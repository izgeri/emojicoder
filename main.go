package main

import (
	"mime"
	"os"
//	"unicode"
	"path"
//	"bytes"

	"errors"
	"io/ioutil"
)

// A tool to transform text into emojis or emojis into text
// Usage:
//   ./emojicode [encode/decode] [input file] [output file]
func main() {

	if len(os.Args) < 2 {
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

	emojiString := stringToEmoji(encodedFile)

	panic(errors.New(emojiString))
//	fmt.Printf("emojied file: %v", emojiString)
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

/*
  str := "01234567890abcdefghijklmnopqrstuvwxyz??\n01234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ";
  emo := stringToEmoji(str);
*/

func stringToEmoji(str string) string {
  runes := []rune(str)
  outrunes := make([]rune, len(runes))
  for i := 0; i < len(runes); i = i+1 {
    outrunes[i] = translateRune(runes[i])
  }
  return string(outrunes)
}

func translateRune(r rune) rune {
  var lowerA = '😄'
  if (r >= 'a' && r <= 'z') {
    return lowerA + (r - 'a');
  }
  var capA = '😤'
  if (r >= 'A' && r <= 'Z') {
    return capA + (r - 'A');
  }
  var num = '🤐'
  if (r >= '0' && r <= '9') {
    return num + (r - '0');
  }
  if (r == '?') {
    return '💥'
  }
  return r
}
/*
// reads in a file, 
readAndDecodeFile(fileName string) string {

}
*/
