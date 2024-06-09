package autocorrect

import (
	"strconv"
)

func Hex(input string) string {
	if len(input) == 0 {
		return ""
	}

	base := 1
	num := 0
	answer := ""

	var outTemp string = input
	var startpattern string = ""
	var endpattern string = ""

	//check for punctuations at the end
	for i := len(outTemp) - 1; i > 0 && (outTemp[i] == '(' || outTemp[i] == ')'); i-- {
		endpattern = string(input[i:])
		outTemp = input[:i]
	}

	//check for punctuations at the beginning
	for i := 0; i < len(outTemp) && (input[i] == '(' || input[i] == ')'); i++ {
		startpattern += string(input[i])
		outTemp = outTemp[1:]
	}

	//confirm that the string only has valid characters like 0-9 or A-F or a-f
	for i := 0; i < len(outTemp); i++ {
		digit := outTemp[i]
		if (digit < '0' || digit > '9') && (digit < 'A' || digit > 'F') && (digit < 'a' || digit > 'f') {
			//return the input unchanged if it's not a hex number
			return input
		}
	}

	//convert from hex to decimal
	for i := len(outTemp) - 1; i >= 0; i-- {
		digit := outTemp[i]
		var value int
		switch {
		case digit >= '0' && digit <= '9':
			value = int(digit - '0')
		case digit >= 'A' && digit <= 'F':
			value = int(digit - 'A' + 10)
		case digit >= 'a' && digit <= 'f':
			value = int(digit - 'a' + 10)
		}
		num += value * base
		base *= 16
	}

	answer = strconv.FormatInt(int64(num), 10)

	//return it as a string
	return startpattern + answer + endpattern
}
