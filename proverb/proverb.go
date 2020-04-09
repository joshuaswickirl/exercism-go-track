package proverb

import "fmt"

// Proverb returns the saying based on the input.
func Proverb(rhyme []string) (proverb []string) {
	for i := 0; i < len(rhyme); i++ {
		var phrase string
		if i == len(rhyme)-1 {
			phrase = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		} else {
			phrase = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		}
		proverb = append(proverb, phrase)
	}
	return
}
