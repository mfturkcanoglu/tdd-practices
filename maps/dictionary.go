package main

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound         = errors.New("could not find the word you were looking for")
	ErrWordExists       = errors.New("word is already exists in dictionary")
	ErrWordDoesNotExist = errors.New("cannot update word because it does not exist")
)

func (d Dictionary) Search(keyword string) (string, error) {
	definition, ok := d[keyword]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(keyword, value string) error {
	_, ok := d[keyword]
	if ok {
		return ErrWordExists
	}
	d[keyword] = value
	return nil
}

func (d Dictionary) Update(keyword, newValue string) error {
	_, err := d.Search(keyword)
	switch err {
	case nil:
		d[keyword] = newValue
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
