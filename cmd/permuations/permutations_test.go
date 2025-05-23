package permutations

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	t.Run("should return ['a'] when input is a", func(t *testing.T) {
		result := Permutations("a")

		if len(result) != 1 && result[0] != "a" {
			t.Error("Expected ['a'], but got", result)
		}
	})

	t.Run("should return permution of ab", func(t *testing.T) {
		result := Permutations("ab")

		if !reflect.DeepEqual(result, []string{"ab", "ba"}) {
			t.Error("Expected permution of ab, but got", result)
		}
	})

	t.Run("should return permution of abc", func(t *testing.T) {
		result := Permutations("abc")
		expected := []string{"abc", "acb", "bac", "bca", "cab", "cba"}

		if !reflect.DeepEqual(result, expected) {
			t.Error("Expected permution of abc, but got", result)

		}
	})
}
