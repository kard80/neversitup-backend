package smily

import "testing"

func TestCountSmileys(t *testing.T) {
	t.Run("should throw panic when input range is not 2 or 3", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic, but did not get it")
			}
		}()
		CountSmileys([]string{":)", ":-D/"})
	})

	t.Run("should return 2", func(t *testing.T) {
		input := []string{":)", ";(", ";}", ":-D"}
		result := CountSmileys(input)
		if result != 2 {
			t.Error("Expected 2, but got", result)
		}
	})
	t.Run("should return 3", func(t *testing.T) {
		input := []string{";D", ":-(", ":-)", ";~)"}
		result := CountSmileys(input)
		if result != 3 {
			t.Error("Expected 3, but got", result)
		}
	})

	t.Run("should return 1", func(t *testing.T) {
		input := []string{";]", ":[", ";*", ":$", ";-D"}
		result := CountSmileys(input)
		if result != 1 {
			t.Error("Expected 1, but got", result)
		}
	})
}

func TestHandleSmiley(t *testing.T) {
	t.Run("should return false when eye is invalid", func(t *testing.T) {
		input := ".)"
		result := handleSmiley(input)
		if result != false {
			t.Error("Expected false, but got", result)
		}
	})

	t.Run("should return false when mounth is invalid", func(t *testing.T) {
		input := ":("
		result := handleSmiley(input)
		if result != false {
			t.Error("Expected false, but got", result)
		}
	})

	t.Run("should return true when eye and mount is valid", func(t *testing.T) {
		input := ":)"
		result := handleSmiley(input)
		if result != true {
			t.Error("Expected true, but got", result)
		}
	})
}

func TestHandleSmileyWithNose(t *testing.T) {
	t.Run("should return false when eye is invalid", func(t *testing.T) {
		input := ".-)"
		result := handleSmileyWithNose(input)
		if result != false {
			t.Error("Expected false, but got", result)
		}
	})

	t.Run("should return false when nose is invalid", func(t *testing.T) {
		input := ":|)"
		result := handleSmileyWithNose(input)
		if result != false {
			t.Error("Expected false, but got", result)
		}
	})

	t.Run("should return false when mounth is invalid", func(t *testing.T) {
		input := ":-("
		result := handleSmileyWithNose(input)
		if result != false {
			t.Error("Expected false, but got", result)
		}
	})

	t.Run("should return true when eye, nose, and mounth are valid", func(t *testing.T) {
		input := ":-)"
		result := handleSmileyWithNose(input)
		if result != true {
			t.Error("Expected false, but got", result)
		}
	})
}
func TestIsValidEyes(t *testing.T) {
	t.Run("should return true when eye is :", func(t *testing.T) {
		input := ":"
		result := isValidEyes(input)
		if result != true {
			t.Error("Expected true, but got", result)
		}
	})

	t.Run("should return true when eye is ;", func(t *testing.T) {
		input := ";"
		result := isValidEyes(input)
		if result != true {
			t.Error("Expected true, but got", result)
		}
	})

	t.Run("should return false when eye is not : or ;", func(t *testing.T) {
		input := "|"
		result := isValidEyes(input)
		if result != false {
			t.Error("Expected false, but got", result)
		}
	})
}
func TestIsValidNose(t *testing.T) {
	t.Run("should return true when nose is -", func(t *testing.T) {
		input := "-"
		result := isValidNose(input)
		if result != true {
			t.Error("Expected true, but got", result)
		}
	})

	t.Run("should return true when nose is ~", func(t *testing.T) {
		input := "-"
		result := isValidNose(input)
		if result != true {
			t.Error("Expected true, but got", result)
		}
	})

	t.Run("should return true when nose is not - or ~", func(t *testing.T) {
		input := "/"
		result := isValidNose(input)
		if result != false {
			t.Error("Expected false, but got", result)
		}
	})
}
func TestIsValidMounth(t *testing.T) {
	t.Run("should return true when mounth is )", func(t *testing.T) {
		input := ")"
		result := isValidMounth(input)
		if result != true {
			t.Error("Expected true, but got", result)
		}
	})

	t.Run("should return true when mounth is D", func(t *testing.T) {
		input := "D"
		result := isValidMounth(input)
		if result != true {
			t.Error("Expected true, but got", result)
		}
	})

	t.Run("should return false when mounth is not ) or D", func(t *testing.T) {
		input := "d"
		result := isValidMounth(input)
		if result != false {
			t.Error("Expected false, but got", result)
		}
	})
}
