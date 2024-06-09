package autocorrect

func Up(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			result += string(s[i] - 32)
		} else {
			result += string(s[i])
		}
	}
	return result
}
