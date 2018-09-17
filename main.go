package main

import (
	"mime"
	"os"
//	"unicode"
	"path"
//	"bytes"

	"fmt"
	"errors"
	"io/ioutil"
)

func main() {

	if len(os.Args) == 1 {
		panic(errors.New("Must include file"))
	}

	file := os.Args[1]

	if !path.IsAbs(file) {
		panic(errors.New("Must pass absolute path"))
	}

	if _, err := os.Stat(file); os.IsNotExist(err) {
		panic(errors.New("File does not exist!"))
	}

	encodedFile := readAndEncodeFile(file)

	fmt.Printf("encoded file: %v", encodedFile)
}

// reads in a file, base 64 encodes it, and returns a string
func readAndEncodeFile(fileName string) string {

	var encodedFile string

	// Read file
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)

	// Encode file
	encodedFile = mime.BEncoding.Encode("UTF-8", fileString)

	return encodedFile
}
