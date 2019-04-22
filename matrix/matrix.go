package matrix

import (
  "strings"
  "strconv"
  "math"
  "errors"
)

type matrix [][]int

func New(input string) (matrix, error) {
  var m matrix
  rows := strings.Split(input, "\n")
  for i := 0; i < len(rows); i++ {
    row := strings.Split(strings.TrimSpace(rows[i]), " ")
    intRow, err := stringArrayToIntArray(row)
    if err != nil {
      return nil, err
    }
    if i >= 1 && len(m[i-1]) != len(intRow) {
      return nil, errors.New("Invalid Input")
    }
    m = append(m, intRow)
  }
  return  m, nil
}

func stringArrayToIntArray(strs []string) ([]int, error) {
  arr := make([]int, len(strs))
  for i := range arr {
     temp, err := strconv.Atoi(strs[i])
    if temp >= math.MaxInt64 || err != nil {
      return arr, errors.New("Invalid Input")
    }
    arr[i] = temp
  }
  return arr, nil
}

func (m matrix) Rows() [][]int {
  rows := make([][]int, len(m))
  for r, row := range m {
    rows[r] = make([]int, len(row))
    for c, val := range row {
      rows[r][c] = val
    }
  }
  return rows
}

func (m matrix) Cols() [][]int {
  cols := make([][]int, len(m[0]))
  for i := 0; i < len(cols); i++ {
    cols[i] = make([]int, len(m))
  }
  for c, row := range m {
    for r, val := range row {
      cols[r][c] = val
    }
  } 
  return cols
}

func (m matrix) Set(r int, c int, val int) bool {
  if r < 0 || c < 0 || r >= len(m) || c >= len(m) {
    return false
  }
  m[r][c] = val
  return true
}