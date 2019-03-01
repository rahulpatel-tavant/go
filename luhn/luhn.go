package luhn

import(
  "strings"
  "strconv"
) 

func Valid(input1 string) bool {
  input := strings.Replace(input1, " ", "", -1)
  if(len(input) <= 1){
    return false
  }

  secondDigitFlag := true
  sum := 0

  for i := len(input)-1; i >= 0; i-- {
    ascii := int(input[i])
    if ascii < 48 || ascii > 57 {
      return false
    }
    val, _ := strconv.Atoi(string(input[i]))
    if secondDigitFlag == true {
      secondDigitFlag = false
      sum += val
    } else {
      secondDigitFlag = true
      double := val * 2
      if double > 9 {
        sum += (double - 9)
      } else {
        sum += double
      }
    }
  }
  return  sum % 10 == 0
}