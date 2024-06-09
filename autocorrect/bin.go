package autocorrect

import (
	"strconv"
)

func Bin(input string) string {
	if len(input) == 0 {
		return ""
	}
	var outTemp string = input
	var startpattern string = ""
	var endpattern string = ""
	var output int = 0

	//check for punstuations at the end
	for i := len(outTemp) - 1; i > 0 && (outTemp[i] == '(' || outTemp[i] == ')'); i-- {
		endpattern = string(input[i:])
		outTemp = input[:i]
	}

	//check for punctuations at the beginning
	for i := 0; i < len(outTemp) && (input[i] == '(' || input[i] == ')'); i++ {
		startpattern += string(input[i])
		outTemp = outTemp[1:]

	}

	//check that the string only has 0s and 1s
	for i := 0; i < len(outTemp); i++ {
		if outTemp[i] != '0' && outTemp[i] != '1' {
			return input
			//returns the input unchanged if it's not a binary number
		}
	}

	temp, err := strconv.Atoi(outTemp)
	if err != nil {
		return ""
	}

	//convert it to decimal base as an integer
	for i, power := len(outTemp), 1; i > 0; i, power = i-1, power*2 {
		output = output + ((temp % 10) * power)
		temp /= 10
	}

	//add the start and end punctuations back
	//return it as a string
	return startpattern + strconv.Itoa(output) + endpattern
}
