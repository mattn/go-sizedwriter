package sizedwriter

import (
	"errors"
	"io"
	"os"
)

type cb func(*Writer) error

type Writer struct {
	Filename string
	Size     int64
	Perm     os.FileMode
	Cb       cb
	file     *os.File
}

func NewWriter(filename string, size int64, perm os.FileMode, over cb) io.WriteCloser {
	return &Writer{filename, size, perm, over, nil}
}

func (sw *Writer) Write(b []byte) (int, error) {
	fi, err := os.Lstat(sw.Filename)
	var size int64
	if err == nil {
		size = fi.Size() + int64(len(b))
	} else if os.IsNotExist(err) {
		size = int64(len(b))
	}
	if size > sw.Size {
		if sw.Cb != nil {
			sw.file.Close()
			sw.file = nil
			err = sw.Cb(sw)
			if err != nil {
				return 0, err
			}
		} else {
			return 0, errors.New("Can't write more")
		}
	}

	if sw.file == nil {
		sw.file, err = os.OpenFile(sw.Filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, sw.Perm)
		if err != nil {
			return 0, err
		}
	}

	n, err := sw.file.Write(b)
	sw.file.Sync()
	return n, err
}

func (sw *Writer) Close() error {
	var err error
	if sw.file != nil {
		err = sw.file.Close()
		if err == nil {
			sw.file = nil
		}
	}
	return err
}
