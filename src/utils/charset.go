package utils

type CharacterSet struct {
	Lowercase bool
	Uppercase bool
	Numbers   bool
	Symbols   bool
}

func CharsetToString(charset CharacterSet) string { // (string, string) {
	output := ""
	// human := ""
	if charset.Lowercase {
		output += "abcdefghijklmnopqrstuvwxyz"
		// human += "Lowercase, "
	}
	if charset.Uppercase {
		output += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		// human += "Uppercase, "
	}
	if charset.Numbers {
		output += "0123456789"
		// human += "Numbers, "
	}
	if charset.Symbols {
		output += "!@#$%^&*"
		// human += "Symbols, "
	}

	// if len(human) > 0 {
	// 	human = human[:len(human)-2]
	// }

	return output // , human
}

func RemoveDuplicateLetters(s string) string {
	seen := make(map[rune]bool)
	result := make([]rune, 0, len(s))

	for _, char := range s {
		if _, exists := seen[char]; !exists {
			seen[char] = true
			result = append(result, char)
		}
	}

	return string(result)
}
