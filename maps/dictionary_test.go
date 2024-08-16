package dictionary

import (
	"testing"
)

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertValue(t testing.TB, d Dictionary, key, want string) {
	t.Helper()

	got, err := d.Search(key)

	if err != nil {
		t.Fatal("did not expect error but got error")
	}

	assertStrings(t, got, want)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("testing key that exists", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("testing key that does not exist", func(t *testing.T) {
		_, err := dictionary.Search("doesNotExist")

		if err == nil {
			t.Fatal("expected an error, but did not get error")
		}

		want := ErrKeyNotFound

		assertError(t, err, want)
	})

}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is to test adding to dictionary"

		dictionary.Add(key, value)

		assertValue(t, dictionary, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "adding"

		dictionary := Dictionary{key: value}

		exists := "already exists"

		err := dictionary.Add(key, exists)

		if err == nil {
			t.Fatal("expected an error, but did not get error")
		}

		want := ErrKeyInUse

		if err != nil {
			assertError(t, err, want)
		}
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update existing key-value pair", func(t *testing.T) {
		key := "test"
		existingValue := "this is just a test"
		dictionary := Dictionary{key: existingValue}
		newValue := "this is the new value to update"
		err := dictionary.Update(key, newValue)

		assertError(t, err, nil)
		assertValue(t, dictionary, key, newValue)

	})

	t.Run("trying to update non-existent key", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "newKey"
		value := "newValue"
		err := dictionary.Update(key, value)

		assertError(t, err, ErrNoKeyToUpdate)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Deleting key-value pair in map", func(t *testing.T) {
		key := "test"
		value := "to delete"
		dictionary := Dictionary{key: value}
		dictionary.Delete(key)

		_, err := dictionary.Search(key)
		assertError(t, err, ErrKeyNotFound)
	})
}
