# Context variables.
ifeq (${VERSION},)
# populate version if not available as env var
    VERSION := $(shell git rev-parse --short HEAD)
endif
UNAME_S := $(shell uname -s)
APP_NAME := capelo
GH_USERNAME := luiscape

# Colors for makin things prettier.
magenta="\\033[34m"
green="\\033[32m"
yellow="\\033[33m"
cyan="\\033[36m"
white="\\033[37m"
reset="\\033[0m"

# Output help text for each command.
# Reference: https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## execute this help command
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)


#####################################################
#    ______   _______  __   __  _______  ___        #
#   |      | |       ||  | |  ||       ||   |       #
#   |  _    ||    ___||  |_|  ||    ___||   |       #
#   | | |   ||   |___ |       ||   |___ |   |       #
#   | |_|   ||    ___||       ||    ___||   |___    #
#   |       ||   |___  |     | |   |___ |       |   #
#   |______| |_______|  |___|  |_______||_______|   #
#                                                   #
#####################################################

version: ## prints version and exits
	@echo " → capelo version: ${green}$(VERSION)${reset}"

develop: ## setup python local virtual env
	@printf "${white}\n> ${magenta}creating local development environment${reset} \n"

	# download go dependencies
	go mod download;

	# test formatting
	go install github.com/mfridman/tparse@latest;

	@printf "${white}\n> ${magenta}done! go build something ${green}:)${reset}\n"

test: ## runs all tests with py.test and reports coverage
	@printf "${white} \n> ${magenta}running test suite${reset} \n";
	go test -cover -json -timeout 20m ./... | tee test.json | tparse -smallscreen -top;
	rm test.json;


##############################################################
#    ______   _______  __   __  _______  _______  _______    #
#   |      | |       ||  | |  ||       ||       ||       |   #
#   |  _    ||    ___||  |_|  ||   _   ||    _  ||  _____|   #
#   | | |   ||   |___ |       ||  | |  ||   |_| || |_____    #
#   | |_|   ||    ___||       ||  |_|  ||    ___||_____  |   #
#   |       ||   |___  |     | |       ||   |     _____| |   #
#   |______| |_______|  |___|  |_______||___|    |_______|   #
#                                                            #
##############################################################

build: check-env build-frontend
build:
	@echo " → building ${green}tools${reset}\n";
	go build -o ${APP_NAME}
	@echo "binary built successfully: ./${APP_NAME}"

build-frontend: ## builds the frontend
	@echo " → building ${green}frontend${reset}\n";
	# build ui bundle outside of container
	cd ui && yarn build && mv -v dist/* ../${GH_USERNAME}.github.io/;

check-env: ## checks that required env vars are defined
	@echo "no env to check"

deploy: ## updates github pages site
	@echo " → deploying to ${green}${GH_USERNAME}.github.io${reset}\n";
	cd ${GH_USERNAME}.github.io && git add . && git commit -m "update" && git push origin;
