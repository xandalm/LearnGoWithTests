package main

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dict.Add(word, definition)

		assertError(t, err, nil)
		asserDefinition(t, dict, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}

		err := dict.Add(word, definition)

		assertError(t, err, ErrAlreadyExists)
		asserDefinition(t, dict, word, definition)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf(`got "%s", want "%s"`, got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func asserDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}
