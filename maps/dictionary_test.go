package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "test dictionary"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "test dictionary"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		_ = dictionary.Add("new", "value")

		want := "value"
		got, err := dictionary.Search("new")

		assertDefinition(t, err, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "new"
		value := "value"
		dictionary := Dictionary{word: value}

		errAdd := dictionary.Add(word, "existing value")
		got, errSearch := dictionary.Search(word)

		assertError(t, errAdd, ErrWordExists)
		assertDefinition(t, errSearch, got, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing key", func(t *testing.T) {
		key := "key1"
		newValue := "newValue"
		dictionary := Dictionary{key: "value1"}
		dictionary.Update(key, newValue)

		value, _ := dictionary.Search(key)
		assertStrings(t, value, newValue)
	})

	t.Run("update non-exist key", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("key", "value")
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	key := "key"
	dictionary := Dictionary{key: "value"}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)
	if err == nil {
		t.Errorf("Expected %q to be deleted", key)
	}
}

var assertStrings = func(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Expected %q, but got %q", want, got)
	}
}

var assertError = func(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Fatalf("Expected %q, but got %q", want, got)
	}
}

var assertDefinition = func(t *testing.T, err error, got, want string) {
	t.Helper()

	if err != nil {
		t.Fatalf("should find added word: %q, err :%v", "new", err)
	}
	if got != want {
		t.Errorf("Expected %q, but got %q", want, got)
	}
}
