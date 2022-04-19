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

	f := os.DirFS(filepath.Join(cwd, *dir))
	fd := fdup.NewFdup(f)

	filesMap, err := fd.Search()
	if err != nil {
		log.Fatal(err)
	}
	fd.PrintDuplicateFiles(filesMap)
}
