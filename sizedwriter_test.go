package sizedwriter_test

import (
	"crypto/rand"
	"encoding/hex"
	sw "github.com/mattn/go-sizedwriter"
	"os"
	"path/filepath"
	"testing"
)

func TempFilename() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return filepath.Join(os.TempDir(), "sizedwriter"+hex.EncodeToString(randBytes)+".log")
}

func TestSimple(t *testing.T) {
	filename := TempFilename()
	defer os.Remove(filename)

	limited := false
	sw := sw.NewWriter(filename, 500, 0644, func(sw *sw.Writer) error {
		limited = true
		return nil
	})

	b := make([]byte, 499)
	n, err := sw.Write(b)
	if err != nil {
		t.Fatal(err)
	}
	if n != 499 {
		t.Fatalf("Loss some bytes")
	}
	if limited {
		t.Fatalf("No expected limited")
	}

	n, err = sw.Write([]byte{0})
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("Loss some bytes")
	}
	if limited {
		t.Fatalf("No expected limited")
	}

	n, err = sw.Write([]byte{0})
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("Loss some bytes")
	}
	if !limited {
		t.Fatalf("Expected limited")
	}
}

func TestLimited(t *testing.T) {
	filename := TempFilename()
	defer os.Remove(filename)

	limited := false
	sw := sw.NewWriter(filename, 500, 0644, func(sw *sw.Writer) error {
		limited = true
		return nil
	})

	b := make([]byte, 501)
	n, err := sw.Write(b)
	if err != nil {
		t.Fatal(err)
	}
	if n != 501 {
		t.Fatalf("Loss some bytes")
	}
	if !limited {
		t.Fatalf("Expected limited")
	}
}
