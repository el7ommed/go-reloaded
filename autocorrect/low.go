package autocorrect

func Low(s string) string {
	result := ""
	if len(s) == 0 {
		return result
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			result += string(s[i] + 32)
		} else {
			result += string(s[i])
		}
	}
	return result
}
