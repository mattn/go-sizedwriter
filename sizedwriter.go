package sizedwriter

import (
	"io"
	"os"
)

type cb func(*SizedWriter) error

type SizedWriter struct {
	Filename string
	Size     int64
	Perm     os.FileMode
	Cb       cb
	file     *os.File
}

func NewSizedWriter(filename string, size int64, perm os.FileMode, over cb) io.WriteCloser {
	return &SizedWriter{filename, size, perm, over, nil}
}

func (sw *SizedWriter) Write(b []byte) (int, error) {
	fi, err := os.Lstat(sw.Filename)
	if err == nil {
		if fi.Size()+int64(len(b)) > sw.Size {
			sw.file.Close()
			sw.file = nil
			err = sw.Cb(sw)
			if err != nil {
				return 0, err
			}
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

func (sw *SizedWriter) Close() error {
	var err error
	if sw.file != nil {
		err = sw.file.Close()
		if err == nil {
			sw.file = nil
		}
	}
	return err
}
