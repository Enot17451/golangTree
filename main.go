package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	out := new(bytes.Buffer)
	dirTree(out, "testdata", true)
	dirTree(out, "testdata", false)
}

func takeAll() {

}

func dirTree(out *bytes.Buffer, path string, onlyDirs bool) error {
	dir, er := os.ReadDir("./" + path)
	if er != nil {
		return er
	}
	for _, record := range dir {
		if record.IsDir() && !onlyDirs {
			fmt.Println("----" + record.Name())
		} else {
			fmt.Println(record.Name())
		}

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
