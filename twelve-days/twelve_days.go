package twelve

import "fmt"

type verse struct {
	day, gift string
}

var verses = [12]verse{
	{"first", "a Partridge in a Pear Tree."},
	{"second", "two Turtle Doves, and "},
	{"third", "three French Hens, "},
	{"fourth", "four Calling Birds, "},
	{"fifth", "five Gold Rings, "},
	{"sixth", "six Geese-a-Laying, "},
	{"seventh", "seven Swans-a-Swimming, "},
	{"eighth", "eight Maids-a-Milking, "},
	{"ninth", "nine Ladies Dancing, "},
	{"tenth", "ten Lords-a-Leaping, "},
	{"eleventh", "eleven Pipers Piping, "},
	{"twelfth", "twelve Drummers Drumming, "},
}

// Verse returns the verse for the given day.
func Verse(number int) string {
	verse := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ",
		verses[number-1].day)
	for i := number - 1; i >= 0; i-- {
		verse += verses[i].gift
	}
	return verse
}

// Song returns all verses in the song.
func Song() string {
	var song string
	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}
	return song[:len(song)-1]
}
