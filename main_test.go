package main

import (
	"testing"
)

func TestParsing(t *testing.T) {

	filepath, filename := parseRawFilename("./path/to/file.yaml")

	if filepath != "./path/to" {
		t.Fatal("FILEPATH IS FUKD: "+filepath)
	}

	if filename != "file.yaml" {
		t.Fatal("FILENAME IS FUKD: "+filename)
	}

}
