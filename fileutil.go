package fileutil

import (
	"bufio"
	"errors"
	"os"
)

const (
	Read int = iota
	Write
	RW
)

type FileUtil struct {
	FileName    string
	RW          int
	Body        []string
	FilePointer *os.File
	Err         error
}

func NewFileUtil(fname string, rw int) *FileUtil {
	s := new(FileUtil)

	s.FileName = fname
	s.RW = rw
	s.Err = nil

	return s
}

func (s *FileUtil) ClearErr() {
	s.errHandle(false)
}

func (s *FileUtil) Open() *FileUtil {
	f, err := os.Open(s.FileName)
	if err != nil {
		s.errHandle(true)("FileOpenError")
	}
	s.FilePointer = f

	return s
}

func (s *FileUtil) Close() {
	s.FilePointer.Close()
}

func (s *FileUtil) GetErr() error {
	return s.Err
}

func (s *FileUtil) errHandle(exist bool) func(e string) {
	if exist {
		return func(e string) {
			s.Err = errors.New(e)
		}
	} else {
		s.Err = nil
		return nil
	}
}

func (s *FileUtil) IsErrFound() bool {
	if s.Err != nil {
		return true
	} else {
		return false
	}
}

func (s *FileUtil) ReadLines() []string {
	lines := make([]string, 0, 200)
	scanner := bufio.NewScanner(s.FilePointer)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		s.errHandle(true)("FileScanError")
	}

	return lines
}
