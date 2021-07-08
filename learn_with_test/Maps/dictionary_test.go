package Maps

import (
	"testing"
)

var thisIsAJustTest = "this is a just test"

// 因为map是引用类型,所以如果直接初始化会得到一个nil异常
// 使用 make 或者 在尾部加上{}可以
//var m map[string]string
//
//
//
//func TestMapNil(t *testing.T){
//	var dictionary = make(map[string]string)
//	if dictionary == nil {
//		t.Errorf("maps is nil not init")
//	}
//}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	newDefinition := "new definition"
	dictionary.Update(word, newDefinition)
	assertDefinition(t, dictionary, word, newDefinition)
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		dictionary.Add(word, definition)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if definition != got {
		t.Errorf("got '%s' want '%s'", got, definition)
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
