package Maps

import (
	"testing"
)

var thisIsAJustTest = "this is a just test"

var m map[string]string

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")
	want := "this is just a test"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if want != got {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": thisIsAJustTest}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := thisIsAJustTest
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
