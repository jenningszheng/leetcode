package goat_latin

func isVowel(letter uint8) bool {
	switch letter {
	case 65, 69, 73, 79, 85, 97, 101, 105, 111, 117:
		return true
	default:
		return false
	}
}

func ToGoatLatin(sentence string) string {
	n := len(sentence)
	i := 0
	j := 0
	k := 1
	for i < n {
		if j >= n || sentence[j] == 32 {
			suffix := "ma"
			if !isVowel(sentence[i]) {
				suffix = string(sentence[i]) + suffix
			}
			for l := 0; l < k; l++ {
				suffix += "a"
			}
			if isVowel(sentence[i]) {
				sentence = sentence[:j] + suffix + sentence[j:]
			} else {
				sentence = sentence[:i] + sentence[i+1:j] + suffix + sentence[j:]
			}
			j += 3 + k
			i = j
			n += 2 + k
			k++
		}
		j++
	}
	return sentence
}
