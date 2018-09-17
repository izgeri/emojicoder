package main

import (
  "fmt"
//  "mime"
)

func main() {
//  str := mime.BEncoding.Encode("utf-8", "¡Hola, señor!")
  str := "01234567890abcdefghijklmnopqrstuvwxyz??\n01234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ";
  fmt.Println(str);
  emo := stringToEmoji(str);
  fmt.Println(emo);
}

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
  fmt.Println(int(lowerA))
  fmt.Println(int(capA))
  fmt.Println(int(num))
  if (r >= '0' && r <= '9') {
    return num + (r - '0');
  }
  if (r == '?') {
    return '💥'
  }
  return r
}
