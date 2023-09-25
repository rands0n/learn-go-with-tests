package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("do a simple sum", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if expected != sum {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}
