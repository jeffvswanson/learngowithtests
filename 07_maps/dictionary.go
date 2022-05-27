package main

import "errors"

type Dictionary map[string]string

var ErrNotFound error = errors.New("word not in dictionary")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
