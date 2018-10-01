package twelve

var verse = map[string]string{
	"first":    "a Partridge in a Pear Tree.",
	"twelfth":  "twelve Drummers Drumming",
	"eleventh": "eleven Pipers Piping",
	"tenth":    "ten Lords-a-Leaping",
	"ninth":    "nine Ladies Dancing",
	"eighth":   "eight Maids-a-Milking",
	"seventh":  "seven Swans-a-Swimming",
	"sixth":    "six Geese-a-Laying",
	"fifth":    "five Gold Rings",
	"fourth":   "four Calling Birds",
	"third":    "three French Hens",
	"second":   "two Turtle Doves",
}

var wording = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

func Verse(i int) string {
	line := "On the "
	line += wording[i]
	line += " day of Christmas my true love gave to me, "
	secondLine := ""
	for ind := i; ind > 1; ind--  {
		key := wording[ind]
		secondLine += verse[key] + ", "
	}
	if i == 1{
		secondLine += "" + verse[wording[1]]
	} else {
		secondLine += "and " + verse[wording[1]]
	}
	line += secondLine
	return line
}

func Song() string {
	var song = ""
	for ind := 1; ind <= len(wording); ind++{
		song += Verse(ind) + "\n"
	}
	return song
}
