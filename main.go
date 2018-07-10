package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var prefix string
var root string

func zipconcat(zipFiles []os.FileInfo) error {
	writers := map[string]*os.File{}

	for _, zipFile := range zipFiles {
		r, err := zip.OpenReader(path.Join(root, zipFile.Name()))
		if err == zip.ErrFormat {
			continue
		}
		if err != nil {
			return fmt.Errorf("could not open archive reader on %s: %v", zipFile.Name(), err)
		}
		defer r.Close()

		for _, f := range r.File {
			if _, ok := writers[f.Name]; !ok {
				outpath := path.Join(root, fmt.Sprintf("%s_%s", prefix, f.Name))
				rw, err := os.Create(outpath)
				if err != nil {
					return fmt.Errorf("could not open output file %s: %v", outpath, err)
				}

				writers[f.Name] = rw
			}

			rc, err := f.Open()
			defer rc.Close()
			if err != nil {
				return fmt.Errorf("could not open file %s in archive %s: %v", f.Name, zipFile.Name(), err)
			}
			_, err = io.Copy(writers[f.Name], rc)
			if err != nil {
				return fmt.Errorf("could not read content of file %s in archive %s: %v", f.Name, zipFile.Name(), err)
			}
		}

	}
	for _, writer := range writers {
		writer.Close()
	}
	return nil
}

func main() {
	flag.StringVar(&prefix, "prefix", "", "prefix to use in output filenames")
	flag.StringVar(&root, "root", "", "path to directory containing zip archives")
	flag.Parse()

	if root == "" {
		flag.Usage()
		os.Exit(0)
	}

	if fi, err := os.Stat(root); err != nil || !fi.Mode().IsDir() {
		log.Fatalf("%s is not a directory", root)
	}

	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatalf("could not list directory %s: %v", root, err)
	}

	err = zipconcat(files)
	if err != nil {
		log.Fatalf("could not zipconcat files in %s: %v", root, err)
	}
}
