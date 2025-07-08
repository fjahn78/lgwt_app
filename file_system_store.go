package main

import "strings"

type FileSystemPlayerStore struct {
	db *strings.Reader
}

func (f FileSystemPlayerStore) GetLeague() []Player {
	return nil
}
