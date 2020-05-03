package encode

import (
	"fmt"
	"strconv"
	"strings"
)

func RunLengthEncode(data string) string {
	var encoded string
	var last rune
	var num int
	for _, curr := range data {
		if last == 0 || last == curr {
			last = curr
			num++
		}
		if last != curr && num == 1 {
			encoded += string(last)
			last = curr
		} else if last != curr && num > 1 {
			encoded += fmt.Sprintf("%v%v", num, string(last))
			last = curr
			num = 1
		}
	}
	if num == 1 {
		encoded += string(last)
	} else if num > 1 {
		encoded += fmt.Sprintf("%v%v", num, string(last))
	}
	return encoded
}

func RunLengthDecode(data string) string {
	var decoded string
	var num string
	for _, r := range data {
		if r >= '1' && r <= '9' {
			num += string(r)
		} else {
			if num == "" {
				num = "1"
			}
			n, _ := strconv.Atoi(num)
			decoded += strings.Repeat(string(r), n)
			num = ""
		}
	}
	return decoded
}
