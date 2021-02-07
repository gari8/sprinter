package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/txtar"
)

type archive txtar.Archive

func main() {
	var (
		pref string
	)
	flag.StringVar(&pref, "strip", "", "string which remove from head of path")
	flag.Parse()

	dir := flag.Arg(0)
	if dir == "" {
		log.Fatal("no such a directory")
	}

	output := flag.Arg(1)
	if output == "" {
		log.Fatal("no such a output path")
	}

	var onion, mvc archive

	err := onion.walkTemplate(dir + "/_onion/", pref + "_onion/")

	if err != nil {
		log.Fatal(err)
	}

	err = mvc.walkTemplate(dir + "/_mvc/", pref + "_mvc/")

	if err != nil {
		log.Fatal(err)
	}

	w, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	archivedOnion := archiveToString(onion)
	archivedMVC := archiveToString(mvc)


	if archivedOnion != "" || archivedMVC != "" {
		_, _ = fmt.Fprintln(w, "// DO NOT EDIT.")
		_, _ = fmt.Fprintln(w, "")
		_, _ = fmt.Fprintln(w, "package main")
		_, _ = fmt.Fprintln(w)
		_, _ = fmt.Fprintln(w, `import "text/template"`)
		_, _ = fmt.Fprintln(w)
		_, _ = fmt.Fprintf(w, "var tmplOnion = template.Must(template.New"+
			"(\"template\").Delims(`@@`, `@@`).Parse(%q))\n", archivedOnion)
		_, _ = fmt.Fprintf(w, "var tmplMVC = template.Must(template.New"+
			"(\"template\").Delims(`@@`, `@@`).Parse(%q))\n", archivedMVC)
	}

	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
}

func archiveToString(ar archive) string {
	arc := txtar.Archive(ar)
	return string(txtar.Format(&arc))
}

func (ar *archive)walkTemplate(dir string, pref string) error {
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		base := filepath.Base(path)

		if info.IsDir() {
			if len(base) > 0 && base[0] == '.' {
				return filepath.SkipDir
			}
			return nil
		}

		if len(base) > 0 && base[0] == '.' {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		p := filepath.ToSlash(path)
		ar.Files = append(ar.Files, txtar.File{
			Name: strings.TrimPrefix(p, pref),
			Data: data,
		})
		return nil
	}); err != nil {
		return err
	}
	return nil
}
