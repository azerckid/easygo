package mydict

import (
	"errors"
)

type Dictionary map[string]string

var errNotFound = errors.New("not Found")
var errCantUpdate = errors.New("cant update non-existing word")
var errWordExist = errors.New("that word already exists")
var errNothingDelete = errors.New("nothing to delete")

func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if exist {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExist
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errNothingDelete
	case nil:
		delete(d, word)
	}
	return nil

}
