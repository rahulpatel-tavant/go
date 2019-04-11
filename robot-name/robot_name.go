package robotname

import (
  "math/rand"
  "time"
)

type Robot struct {  
  name string
}

func RandStringBytes(n int, letterBytes string) string {
  b := make([]byte, n)
  for i := range b {
      b[i] = letterBytes[rand.Intn(len(letterBytes))]
  }
  return string(b)
}

func getRobotName() string {
  rand.Seed(time.Now().UTC().UnixNano())
  return RandStringBytes(2, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") + RandStringBytes(3, "0123456789")
}

func (r *Robot) Name() (string, error) {
  if r.name == "" {
    r.Reset()
  }
  return r.name, nil
}

func (r *Robot) Reset() {
  r.name = getRobotName()
}