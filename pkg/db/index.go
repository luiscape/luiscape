package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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
}

type EntryCollection struct {
	Entries []Entry `json:"Entries"`
}

// walks a path structuring files into a database
func Walk(path string) (EntryCollection, error) {

	entries := []Entry{}
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// parent directory
			parent := filepath.Base(filepath.Dir(path))

			// check if path is directory
			entryType := Section
			if info.IsDir() {

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

					// get timestamp if this is a blog entry
					modifiedtime := info.ModTime()

					// add record to slice
					filename := filepath.Join(parent, filepath.Base(path))
					entries = append(entries, Entry{
						Type:         entryType,
						Path:         filename,
						SectionName:  parent,
						CreationTime: modifiedtime,
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

	fmt.Println("scanned database")
	fmt.Printf("%+v\n", entries)
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
