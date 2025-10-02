package main

// pass a map to a function/method, you are indeed copying it, but just the
// pointer part, not the underlying data structure that contains the data.
type Dictionary map[string]string

// func Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }

// ** NEVER initialize a nil map like: var m map[string]string
// it causes runtime panic
// Instead use: var dictionary = map[string]string{}

// OR

// var dictionary = make(map[string]string)

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
