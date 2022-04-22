package main

import (
	"flag"
	"log"
	"os"

	"github.com/enkeyz/go-fdup/internal/fdup"
)

func main() {
	dir := flag.String("d", ".", "full path to the directory")
	flag.Parse()

	fi, err := os.Stat(*dir)
	if err != nil {
		log.Fatal(err)
	}
	if !fi.IsDir() {
		log.Fatalf("%s is not a directory", *dir)
	}

	f := os.DirFS(*dir)
	fd := fdup.NewFdup(f)

	res, err := fd.Search()
	if err != nil {
		log.Fatal(err)
	}

	fdup.PrintResult(os.Stdout, res)
}
