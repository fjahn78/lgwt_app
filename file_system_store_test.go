package main

import (
	"strings"
	"testing"
)

func TestFileSystemPlayerStore(t *testing.T) {
	database := strings.NewReader(`[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)
	t.Run("league from a reader", func(t *testing.T) {

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)

	})
	t.Run("get player score", func(t *testing.T) {
		store := FileSystemPlayerStore{database: database}

		got := store.GetPlayerScore("Chris")

		want := 33

		assertPlayerScore(t, got, want)
	})
}
