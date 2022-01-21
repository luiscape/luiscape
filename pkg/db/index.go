package db

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// looks for files that have this extension in order to make
// a post
const POST_EXTENSION = ".md"
const DEFAULT_DATABASE_FILENAME = "database.json"

// object types for entries
type EntryType string

const (
	Section EntryType = "section"
	Post    EntryType = "post"
)

// entries are organized with their entire path
// and type
type Entry struct {
	Type         EntryType `json:"Type"`
	Path         string    `json:"Path"`
	SectionName  string    `json:"SectionName"`
	CreationTime time.Time `json:"CreationTime"`
	Title        string    `json:"Title"`
}

type EntryCollection struct {
	Entries []Entry `json:"Entries"`
}

// extracts title from a markdown file
// title is considered the first line
// without the header characters
func ExtractTitle(path string) (string, error) {

	// read file
	fileIO, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		return "", err
	}
	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO)
	if err != nil {
		return "", err
	}

	// find title and remove markdown characters
	lines := strings.Split(string(rawBytes), "\n")
	title := lines[0]
	title = strings.ReplaceAll(title, "# ", "")

	return title, nil
}

// walks a path structuring files into a database
func Walk(targetPath string) (EntryCollection, error) {

	entries := []Entry{}
	targetPathParent := filepath.Base(filepath.Dir(targetPath))
	err := filepath.Walk(targetPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// parent directory
			parent := filepath.Base(filepath.Dir(path))

			// check if path is directory
			entryType := Section
			if info.IsDir() {

				// skip indexing root directory as a section
				if parent == targetPathParent {
					return nil
				}

				// add record to slice
				entries = append(entries, Entry{
					Type: entryType,
					Path: parent,
				})

			} else {

				// check if extension is expected file extension
				fileExtension := filepath.Ext(path)
				if fileExtension == POST_EXTENSION {
					entryType = Post

					// extract title from markdown file
					title, err := ExtractTitle(path)
					if err != nil {
						return err
					}

					// get timestamp if this is a blog entry
					modifiedtime := info.ModTime()

					// add record to slice
					filename := filepath.Join(parent, filepath.Base(path))
					entries = append(entries, Entry{
						Type:         entryType,
						Path:         filename,
						SectionName:  parent,
						CreationTime: modifiedtime,
						Title:        title,
					})
				}

			}
			return nil
		})

	entryCollection := EntryCollection{Entries: entries}

	// handle errors walking
	if err != nil {
		return entryCollection, err
	}

	return entryCollection, nil

}

// writes database to json file on disk
func Write(path string, entries EntryCollection) error {
	file, _ := json.MarshalIndent(entries, "", " ")
	err := ioutil.WriteFile(path, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
