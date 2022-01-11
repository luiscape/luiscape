package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestWrite(t *testing.T) {

	// creates temp dir structure
	tempDir, _ := os.MkdirTemp("dir", "walktest")
	tempPath := path.Join(tempDir, DEFAULT_DATABASE_FILENAME)
	defer os.RemoveAll(tempDir)

	entries := EntryCollection{
		Entries: []Entry{
			{
				Type: Section,
				Path: tempDir,
			},
			{
				Type:        Post,
				SectionName: "test",
				Path:        path.Join(tempDir, "test.md"),
			},
		},
	}

	// test that we can write database succesfully
	err := Write(tempPath, entries)
	if err != nil {
		t.Logf("failed to write database")
		t.Fail()
	}

	// test that database can be read correctly
	jsonFile, err := os.Open(tempPath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var readEntries EntryCollection
	json.Unmarshal(byteValue, &readEntries)

	// test db size is the same
	if len(readEntries.Entries) != len(entries.Entries) {
		t.Logf("number of entries in database is wrong")
		t.Fail()
	}

	// test records are the same
	if entries.Entries[0].Type != Section {
		t.Logf("entry order type wrong")
		t.Fail()
	}
}

func TestWalk(t *testing.T) {

	// creates temp dir structure
	tempDir, _ := os.MkdirTemp(os.TempDir(), "walktest")
	defer os.RemoveAll(tempDir)

	// create test file
	tempPath := path.Join(tempDir, "test.md")
	emptyFile, err := os.Create(tempPath)
	if err != nil {
		t.Logf("failed to create test file")
		t.Fail()
	}
	defer emptyFile.Close()

	// create a test dir
	os.Mkdir(path.Join(tempDir, "test-dir"), 0755)
	defer os.Remove(tempPath)

	entries, err := Walk(tempDir)
	if err != nil {
		t.Logf("failed to walk dir")
		fmt.Println(err)
		t.Fail()
	}

	// test database structure is what's expected
	if len(entries.Entries) != 3 {
		t.Logf("entry database wrong")
		fmt.Println(tempDir)
		fmt.Printf("%+v\n", entries)
		t.Fail()
	}
}
