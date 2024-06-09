package autocorrect

func Aan(article string, word string) string {
	if len(word) < 1 {
		return article
	}

	var wordTemp string = word
	var newArticle string = article

	if article == "aN" {
		newArticle = "an"
	}

	//remove punctuations at the beginning of the word
	for i := 0; i < len(wordTemp); i++ {
		if wordTemp[i] == '(' || wordTemp[i] == ')' {
			wordTemp = word[i+1:]
		}
	}

	//check the first letter of the word and return it if it's already correct
	if Low(string(wordTemp[0])) == "a" || Low(string(wordTemp[0])) == "e" || Low(string(wordTemp[0])) == "i" || Low(string(wordTemp[0])) == "o" || Low(string(wordTemp[0])) == "u" || Low(string(wordTemp[0])) == "h" {
		if Low(article) == "an" {
			return newArticle
		}
	} else {
		if Low(article) == "a" {
			return newArticle
		}
	}

	//check the case of the old article
	if Low(string(wordTemp[0])) == "a" || Low(string(wordTemp[0])) == "e" || Low(string(wordTemp[0])) == "i" || Low(string(wordTemp[0])) == "o" || Low(string(wordTemp[0])) == "u" || Low(string(wordTemp[0])) == "h" {
		if article == "A" {
			return "An"
		} else if article == "a" {
			return "an"
		}
	} else {
		if article == "An" || article == "AN" {
			return "A"
		} else if article == "an" {
			return "a"
		}
	}

	// return new article
	// cases:
	// a
	// A
	// an
	// An
	// AN
	return article
}
