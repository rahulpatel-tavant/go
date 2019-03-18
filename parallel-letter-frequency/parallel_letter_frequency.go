package letter

func ConcurrentFrequency(list []string) FreqMap {
  output := FreqMap{}

  c := make(chan FreqMap)
  for _, value := range list {
     go CalculateFrequency(value, c)
  }

  for i := 0; i < len(list); i++ {
    for letter, count := range <-c {
      output[letter] += count
    }
  }
  return output
}


func CalculateFrequency(s string, chnl chan FreqMap) {
  chnl <- Frequency(s)
}