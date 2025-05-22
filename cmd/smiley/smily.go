package smily

import (
	"slices"
)

func CountSmileys(inputs []string) int {
	result := 0
	for _, v := range inputs {
		isSmiley := false
		if len(v) == 2 {
			isSmiley = handleSmiley(v)
		} else if len(v) == 3 {
			isSmiley = handleSmileyWithNose(v)
		} else {
			panic("invalid input length")
		}

		if isSmiley {
			result++
		}

	}

	return result
}

// input always be 2 characters
func handleSmiley(input string) bool {
	eyes := input[0:1]
	mount := input[1:2]

	return isValidEyes(eyes) && isValidMounth(mount)
}

// input always be 3 characters
func handleSmileyWithNose(input string) bool {
	eyes := input[0:1]
	nose := input[1:2]
	mounth := input[2:3]

	return isValidEyes(eyes) && isValidNose(nose) && isValidMounth(mounth)
}

func isValidEyes(input string) bool {
	eyes := []string{":", ";"}
	return slices.Contains(eyes, input)
}
func isValidNose(input string) bool {
	noses := []string{"-", "~"}
	return slices.Contains(noses, input)
}
func isValidMounth(input string) bool {
	mounths := []string{")", "D"}
	return slices.Contains(mounths, input)
}