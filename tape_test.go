package poker

import (
	"io"
	"testing"
)

func Test_tape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &tape{file: file}

	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
