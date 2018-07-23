package fileutil

import (
	"testing"
)

func TestFileUtil(t *testing.T) {
	s := NewFileUtil("README.md", Read)
	s.Open()
	defer s.Close()

	t.Log(s.ReadLines())
}
