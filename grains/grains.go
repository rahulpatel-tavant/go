package grains

import (
  "errors"
  "math"
)

func Square(input int) (uint64, error) {
  if(input < 1 || input > 64){
      return 0, errors.New("invalid input")
    } else {
      return uint64(math.Pow(2, (float64(input) - 1))), nil
    }
}

func Total() uint64 {
  totalGrains := uint64(0)
  for i := 1; i <= 64; i++ {
    grainsOnSquare, _ := Square(i)
    totalGrains += grainsOnSquare
  }
  return totalGrains
}