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

func main() {

	var (
		pref = flag.String("strip", "", "string which remove from head of path")
	)
	flag.Parse()

	dir := flag.Arg(0)
	if dir == "" {
		log.Fatal("no such a directory")
	}

	output := flag.Arg(1)
	if output == "" {
		log.Fatal("no such a output path")
	}

	var ar txtar.Archive

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
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
			Name: strings.TrimPrefix(p, *pref),
			Data: data,
		})

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	w, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	archived := string(txtar.Format(&ar))
	if archived != "" {
		fmt.Fprintln(w, "// DO NOT EDIT.")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "package main")
		fmt.Fprintln(w)
		fmt.Fprintln(w, `import "text/template"`)
		fmt.Fprintln(w)
		fmt.Fprintf(w, "var tmpl = template.Must(template.New"+
			"(\"template\").Delims(`@@`, `@@`).Parse(%q))\n", archived)
	}

	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
}
