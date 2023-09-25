package maps

import "testing"

func TestMap(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dict := Dict{"test": "this is just a test"}

		got, _ := dict.Search("test")
		want := "this is just a test"

		if got != want {
			t.Errorf("want %q got %q, %q", want, got, dict)
		}
	})

	t.Run("unknown word", func(t *testing.T) {
		dict := Dict{"ops": "ops"}

		_, err := dict.Search("test")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})

	t.Run("add a new key", func(t *testing.T) {
		dict := Dict{}
		dict.Add("test", "this is just a test")

		want := "this is just a test"
		got, err := dict.Search("test")

		if err != nil {
			t.Fatal("should find added word:", err)
		}

		assertError(t, err, nil)
		assertStrings(t, got, want)
	})

	t.Run("add an existing key", func(t *testing.T) {
		dict := Dict{"test": "this is just a test"}

		err := dict.Add("test", "this is just a test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, "test", "this is just a test")
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dict Dict, word, def string) {
	t.Helper()

	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, def)
}
