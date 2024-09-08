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

func dirTree(out *bytes.Buffer, path string, onlyDirs bool) error {
	dir, er := os.ReadDir("./" + path)
	if er != nil {
		return er
	}
	for _, record := range dir {
		if record.IsDir() && !onlyDirs {
			fmt.Println(indent1 + record.Name())
		} else {
			fmt.Println(indent0 + record.Name())
		}
		fmt.Println(indentDown + record.Name())
	}
	return nil
}

func readDir(path string, out *string) error {
	dir, er := os.ReadDir("./" + path)
	if er != nil {
		return er
	}
	for _, record := range dir {
		fmt.Println(record.Name())
	}
	return er
}
