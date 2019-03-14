package clock

import (
  "fmt"
)

type clock struct {  
  hours int
  minutes int
}

func New(hours int, minutes int) clock {
  h, m := FormatTime(hours, minutes)
  return clock {h, m}
}

func (c clock) String() string {
  return fmt.Sprintf("%02d", c.hours) + ":" + fmt.Sprintf("%02d", c.minutes)
}

func (c clock) Add(minutes int) clock {
  return ResetClock(c, minutes)
}

func (c clock) Subtract(minutes int) clock {
  return ResetClock(c, -minutes)
}

func FormatTime(hours int, minutes int) (int, int) {
  totalMinutes := hours * 60 + minutes
  if totalMinutes < 0 {
    f := totalMinutes / 1440
    totalMinutes = (-f + 1) * 1440 + totalMinutes
  }
  totalMinutes = totalMinutes % 1440
  return totalMinutes / 60, totalMinutes % 60
}

func ResetClock(c clock, minutes int) clock {
  totalMinutes := (c.hours * 60 + c.minutes + minutes) % 1440
  h, m := FormatTime(totalMinutes / 60, totalMinutes % 60)
  return clock {h, m}
}
