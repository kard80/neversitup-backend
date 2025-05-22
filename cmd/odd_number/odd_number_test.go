package oddnumber

import "testing"

func TestOddNumber(t *testing.T) {
	t.Run("should return 7", func(t *testing.T) {
		input := []int{7}
		result := OddNumber(input)
		if result != 7 {
			t.Error("Expected 7, but got", result)
		}
	})

	t.Run("should return 0", func(t *testing.T) {
		input := []int{0}
		result := OddNumber(input)
		if result != 0 {
			t.Error("Expected 0, but got", result)
		}
	})

	t.Run("should return 1", func(t *testing.T) {
		input := []int{1, 1, 2}
		result := OddNumber(input)
		if result != 2 {
			t.Error("Expected 2, but got", result)
		}
	})
	t.Run("should return 0", func(t *testing.T) {
		input := []int{0, 1, 0, 1, 0}
		result := OddNumber(input)
		if result != 0 {
			t.Error("Expected 0, but got", result)
		}
	})

	t.Run("should return 4", func(t *testing.T) {
		input := []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}
		result := OddNumber(input)
		if result != 4 {
			t.Error("Expected 4, but got", result)
		}
	})

	t.Run("should panic when there is no odd number of times", func(t *testing.T) {
		input := []int{1, 1}
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic, but got none")
			}
		}()
		OddNumber(input)
	})
}
