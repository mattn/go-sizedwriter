package main

import (
	sw "github.com/mattn/go-sizedwriter"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	filename := "foo.log"
	sw := sw.NewWriter(filename, 500, 0644, func(sw *sw.Writer) error {
		var newname string
		for {
			ext := filepath.Ext(sw.Filename)
			newname = filename[:len(filename)-len(ext)] + "-" + time.Now().Format("20060102150405") + ext
			_, err := os.Lstat(newname)
			if os.IsNotExist(err) {
				break
			}
			time.Sleep(1 * time.Millisecond)
		}
		println("rotate", filename, newname)
		return os.Rename(filename, newname)
	})

	log.SetOutput(sw)
	for {
		log.Println("こんにちわ世界")
		time.Sleep(200 * time.Millisecond)
	}
}
