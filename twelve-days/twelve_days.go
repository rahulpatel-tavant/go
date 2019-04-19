package twelve

var items = []string { "Partridge", "Turtle Doves", "French Hens", "Calling Birds", "Gold Rings", "Geese-a-Laying", "Swans-a-Swimming",
                       "Maids-a-Milking", "Ladies Dancing", "Lords-a-Leaping", "Pipers Piping", "Drummers Drumming" }
var days = []string { "first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth" }
var NumberInWorlds = []string { "a", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve" }

func Song() string {
  song := ""
  for i := 1; i <= 12; i++ {
    verse := Verse(i) + "\n"
    song += verse
  }
  return song
}

func Verse(input int) string {
  verse := "On the " + days[input-1] + " day of Christmas my true love gave to me: "
  for i := input - 1; i >= 0; i-- {
    if i == 0 {
      if input == 1 {
        verse += NumberInWorlds[i] + " " + items[i] + " in a Pear Tree."
      } else {
        verse += "and " + NumberInWorlds[i] + " " + items[i] + " in a Pear Tree."
      }
    } else {
      verse += NumberInWorlds[i] + " " + items[i] + ", "
    }
  }
  return verse
}
