package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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
	Type        EntryType
	Path        string
	SectionName string
}

// walks a path structuring files into a database
func Walk(path string) ([]Entry, error) {

	entries := []Entry{}
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// check if path is directory
			entryType := Section
			if info.IsDir() {

				// add record to slice
				entries = append(entries, Entry{
					Type: entryType,
					Path: path,
				})

			} else {

				// check if extension is expected file extension
				fileExtension := filepath.Ext(path)
				if fileExtension == POST_EXTENSION {

					// add record to slice
					entryType = Post
					sectionName := filepath.Base(filepath.Dir(path))
					entries = append(entries, Entry{
						Type:        entryType,
						Path:        path,
						SectionName: sectionName,
					})
				}

			}
			return nil
		})

	// handle errors walking
	if err != nil {
		return entries, err
	}

	fmt.Println("scanned database")
	fmt.Printf("%+v\n", entries)
	return entries, nil

}

// writes database to json file on disk
func Write(path string, entries []Entry) error {
	file, _ := json.MarshalIndent(entries, "", " ")
	err := ioutil.WriteFile(path, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
