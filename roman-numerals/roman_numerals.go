package romannumerals

import (
	"errors"
	"strings"
)

type number struct {
	arabic int
	roman  string
}

var numbers []number = []number{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(num int) (string, error) {
	var romanNumeral string
	if num < 1 || num > 3000 {
		return "", errors.New("num must be between 1 and 3000")
	}
	for _, n := range numbers {
		if num >= n.arabic {
			romanNumeral += strings.Repeat(n.roman, num/n.arabic)
			num = num % n.arabic
		}
	}
	return romanNumeral, nil
}
