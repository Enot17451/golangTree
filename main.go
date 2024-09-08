package main

import (
	"bytes"
	"fmt"
	"os"
)

const (
	indent0    = "├───"
	indent1    = "│\t"
	indentDown = "└───"
)

func main() {
	out := new(bytes.Buffer)
	dirTree(out, "testdata", true)
	dirTree(out, "testdata", false)
}

func dirTree(out *bytes.Buffer, path string, withFiles bool) error {
	dir, er := os.ReadDir("./" + path)
	if er != nil {
		return er
	}
	for _, record := range dir {
		f := record.Name()
		readDir("/"+f, out)
		if record.IsDir() && !withFiles {

		} else {

		}
	}
	fmt.Println(out)
	return nil
}

func readDir(path string, out *bytes.Buffer) error {
	dir, er := os.ReadDir("./" + path)
	if er != nil {
		return er
	}
	for _, record := range dir {
		fmt.Println(record.Name())
	}
	return er
}
