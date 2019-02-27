package raindrops

import (
  "strconv"
  "sort"
) 

func Convert(num int) string {
  output := ""
  divisors_mapping := map[int]string{
    3: "Pling",
    5: "Plang",
    7: "Plong",
  }

  divisors :=  make([]int, 0)

  for k, _ := range divisors_mapping {
    divisors = append(divisors, k)
  }

  sort.Ints(divisors)

  for _, v := range divisors {
    if num % v == 0 {
      output += divisors_mapping[v]
    }
  }

  if output == "" {
    return strconv.Itoa(num)
  } else {
    return output
  }
}
