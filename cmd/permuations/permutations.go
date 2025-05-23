package permutations

func Permutations(input string) []string {
	result := []string{}
	if len(input) <= 1 {
		return []string{input}
	}

	for i, char := range input {
		fixedChar := string(char)
		remainingChar := input[:i] + input[i+1:]

		for _, perm := range Permutations(remainingChar) {
			result = append(result, fixedChar+perm)
		}
	}

	return result
}
