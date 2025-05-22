package smily

import "testing"

func TestCountSmileys(t * testing.T) {
	t.Run("should throw panic when input range is not 2 or 3", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic, but did not get it")
			}
		}()
		CountSmileys([]string{":)", ":-D/"})
	})
}