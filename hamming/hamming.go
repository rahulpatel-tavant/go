package hamming

import "errors"

func Distance(a, b string) (int, error) {
  hamming_distance := 0
  if len(a) == len(b) {
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
          hamming_distance++
        }
    }
    return hamming_distance, nil
  } else {
    return hamming_distance, errors.New("Length of strings are not same")
  }
}
