package gream

import (
	"log"
	"os"
	"path/filepath"
)

func Path(path string) string {
	return filepath.Join(root(), path)
}

func root() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
