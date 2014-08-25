package main

import (
	sw "github.com/mattn/go-sizedwriter"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func main() {
	filename := "foo.log"
	sw := sw.NewWriter(filename, 500, 0644, func(sw *sw.Writer) error {
		var newname string
		ext := filepath.Ext(sw.Filename)
		base := filename[:len(filename)-len(ext)]
		for {
			newname = base + "-" + time.Now().Format("20060102150405") + ext
			_, err := os.Lstat(newname)
			if os.IsNotExist(err) {
				break
			}
			time.Sleep(1 * time.Millisecond)
		}
		err := os.Rename(filename, newname)
		if err != nil {
			return err
		}
		println("rotate", filename, newname)

		files, err := filepath.Glob(base + "-*" + ext)
		if err != nil {
			return err
		}
		if len(files) > 5 {
			sort.Strings(files)
			println("delete", files[0])
			return os.Remove(files[0])
		}
		return nil
	})

	log.SetOutput(sw)
	for {
		log.Println("こんにちわ世界")
		time.Sleep(200 * time.Millisecond)
	}
}
