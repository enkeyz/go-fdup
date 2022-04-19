package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/enkeyz/go-fdup/internal/fdup"
)

func main() {
	dir := flag.String("d", ".", "directory")
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dirPath := filepath.Join(cwd, *dir)
	fi, err := os.Stat(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	if !fi.IsDir() {
		log.Fatalf("%s is not a directory", *dir)
	}

	f := os.DirFS(dirPath)
	fd := fdup.NewFdup(f)

	res, err := fd.Search()
	if err != nil {
		log.Fatal(err)
	}

	fdup.PrintResult(res)
}
