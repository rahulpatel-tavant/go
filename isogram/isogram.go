package isogram

import "strings"

func IsIsogram(phrase string) bool {

  charCounter := make(map[string]bool)

  for _, char := range strings.ToLower(phrase) {
    if charCounter[string(char)] == true && string(char) != " " && string(char) != "-" {
      return false
    } else {
      charCounter[string(char)] = true
    }
  }
  return true
}