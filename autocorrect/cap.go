package autocorrect

func Cap(s string) string {
	if len(s) == 0 {
		return ""
	}

	var result string = ""
	var outTemp string = s
	var startpattern string = ""

	//check for punctuations at the beginning
	for i := 0; i < len(outTemp) && (s[i] == '(' || s[i] == ')'); i++ {
		startpattern += string(s[i])
		outTemp = s[i+1:]
	}

	//make the first character uppercase to start
	if outTemp[0] >= 'a' && outTemp[0] <= 'z' {
		result += string(outTemp[0] - 32)
	} else {
		result += string(outTemp[0])
	}

	//make the rest lower case
	for i := 1; i < len(outTemp); i++ {
		if outTemp[i] >= 'A' && outTemp[i] <= 'Z' {
			result += string(s[i] + 32)
		} else {
			result += string(outTemp[i])
		}
	}

	return startpattern + result
}
