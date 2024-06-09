package autocorrect

import (
	"strconv"
	"strings"
)

func ModifyLine(line string) string {
	if len(line) == 0 {
		return ""
	}

	//remove all white spaces and replace them with a space to prevent errors in the code
	temp := strings.Fields(line)
	line = strings.Join(temp, " ")

	//keep checking for spaces before punctuation marks
	for strings.Contains(line, " ,") {
		line = strings.Replace(line, " ,", ", ", -1)
	}
	for strings.Contains(line, " .") {
		line = strings.Replace(line, " .", ". ", -1)
	}
	for strings.Contains(line, " !") {
		line = strings.Replace(line, " !", "! ", -1)
	}
	for strings.Contains(line, " ?") {
		line = strings.Replace(line, " ?", "? ", -1)
	}
	for strings.Contains(line, " ;") {
		line = strings.Replace(line, " ;", "; ", -1)
	}
	for strings.Contains(line, " :") {
		line = strings.Replace(line, " :", ": ", -1)
	}

	//add spaces outside brackets to prevent errors (it's ok if there are double spaces at this stage)
	line = strings.Replace(line, "(", " (", -1)
	line = strings.Replace(line, ")", ") ", -1)

	//remove spaces inside brackets to prevent errors
	for strings.Contains(line, "( ") {
		line = strings.Replace(line, "( ", "(", -1)
	}
	for strings.Contains(line, " )") {
		line = strings.Replace(line, " )", ")", -1)
	}

	//add spaces after all punctuation marks (it's ok if there are double spaces at this stage)
	line = strings.Replace(line, ",", ", ", -1)
	line = strings.Replace(line, ".", ". ", -1)
	line = strings.Replace(line, "!", "! ", -1)
	line = strings.Replace(line, "?", "? ", -1)
	line = strings.Replace(line, ";", "; ", -1)
	line = strings.Replace(line, ":", ": ", -1)

	//replace weird quotes with the normal ones
	line = strings.Replace(line, "‘", "'", -1)
	line = strings.Replace(line, "’", "'", -1)

	//add spaces before and after quotes only if it's not within a word such as an apostrophe
	line = strings.Replace(line, " '", " ' ", -1)
	line = strings.Replace(line, "' ", " ' ", -1)

	//split the string again by white spaces
	temp = strings.Fields(line)

	//check for single quotes placement
	var place int = 1
	for i := 0; i < len(temp); i++ {
		if (temp[i] == "'" && i+1 < len(temp)) || (temp[i] == "'" && place%2 == 0 && i == len(temp)-1) {
			if place%2 == 1 && temp[i+1] != "'" {
				temp[i+1] = temp[i] + temp[i+1]
				temp[i] = ""
				place++
			} else if place%2 == 1 && temp[i+1] == "'" {
				temp[i] = temp[i] + temp[i+1]
				temp[i+1] = ""
				place += 2
			} else {
				temp[i-1] = temp[i-1] + temp[i]
				temp[i] = ""
				place++
			}
		}
	}
	line = strings.Join(temp, " ")
	temp = strings.Fields(line)

	//check "a" and "an"
	for i := 0; i < len(temp); i++ {
		if (Low(temp[i]) == "a" || Low(temp[i]) == "an") && i+1 < len(temp) {
			temp[i] = Aan(temp[i], temp[i+1])
		}
	}

	//check for arguments between brackets
	for i := 0; i < len(temp); i++ {
		if Low(temp[i]) == "(bin)" && i > 0 {
			temp[i-1] = Bin(temp[i-1])
			temp[i] = ""
		}
	}
	for i := 0; i < len(temp); i++ {
		if Low(temp[i]) == "(hex)" && i > 0 {
			temp[i-1] = Hex(temp[i-1])
			temp[i] = ""
		}
	}
	for i := 0; i < len(temp); i++ {
		if Low(temp[i]) == "(up)" && i > 0 {
			temp[i-1] = Up(temp[i-1])
			temp[i] = ""
		}
	}
	for i := 0; i < len(temp); i++ {
		if Low(temp[i]) == "(up," && i > 0 {
			if i < len(temp)-1 {
				if (temp[i+1])[len(temp[i+1])-1] == ')' {
					reps, _ := strconv.Atoi((temp[i+1])[0 : len(temp[i+1])-1])
					if reps > i {
						reps = i
					}
					for j := reps; j > 0; j-- {
						temp[i-j] = Up(temp[i-j])
					}
					temp[i], temp[i+1] = "", ""
				}
			}
		}
	}
	for i := 0; i < len(temp); i++ {
		if Low(temp[i]) == "(low)" && i > 0 {
			temp[i-1] = Low(temp[i-1])
			temp[i] = ""
		}
	}
	for i := 0; i < len(temp); i++ {
		if Low(temp[i]) == "(low," && i > 0 {
			if i < len(temp)-1 {
				if (temp[i+1])[len(temp[i+1])-1] == ')' {
					reps, _ := strconv.Atoi((temp[i+1])[0 : len(temp[i+1])-1])
					if reps > i {
						reps = i
					}
					for j := reps; j > 0; j-- {
						temp[i-j] = Low(temp[i-j])
					}
					temp[i], temp[i+1] = "", ""
				}
			}
		}
	}
	for i := 0; i < len(temp); i++ {
		if Low(temp[i]) == "(cap)" && i > 0 {
			temp[i-1] = Cap(temp[i-1])
			temp[i] = ""
		}
	}
	for i := 0; i < len(temp); i++ {
		if Low(temp[i]) == "(cap," && i > 0 {
			if i < len(temp)-1 {
				if (temp[i+1])[len(temp[i+1])-1] == ')' {
					reps, _ := strconv.Atoi((temp[i+1])[0 : len(temp[i+1])-1])
					if reps > i {
						reps = i
					}
					for j := reps; j > 0; j-- {
						temp[i-j] = Cap(temp[i-j])
					}
					temp[i], temp[i+1] = "", ""
				}
			}
		}
	}

	//split the string again to remove double spaces causes by deleting the arguments between brackets
	line = strings.Join(temp, " ")

	//keep checking for spaces before punctuation marks
	for strings.Contains(line, " ,") {
		line = strings.Replace(line, " ,", ", ", -1)
	}
	for strings.Contains(line, " .") {
		line = strings.Replace(line, " .", ". ", -1)
	}
	for strings.Contains(line, " !") {
		line = strings.Replace(line, " !", "! ", -1)
	}
	for strings.Contains(line, " ?") {
		line = strings.Replace(line, " ?", "? ", -1)
	}
	for strings.Contains(line, " ;") {
		line = strings.Replace(line, " ;", "; ", -1)
	}
	for strings.Contains(line, " :") {
		line = strings.Replace(line, " :", ": ", -1)
	}

	//remove white spaces one last time
	temp = strings.Fields(line)
	line = strings.Join(temp, " ")

	//return the line to the main function so it can be written to the output file
	return line
}
