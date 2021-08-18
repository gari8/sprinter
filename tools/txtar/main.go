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

	// archive 作成
	var onion, mvc, clean, minimum, hexagonal, layout archive

	err := layout.walkTemplate(dir + "/common/_layout/", pref + "common/_layout/")
	if err != nil {
		log.Fatal(err)
	}

	err = onion.walkTemplate(dir + "/theme/_onion/", pref + "theme/_onion/")
	if err != nil {
		log.Fatal(err)
	}

	err = mvc.walkTemplate(dir + "/theme/_mvc/", pref + "theme/_mvc/")
	if err != nil {
		log.Fatal(err)
	}

	err = clean.walkTemplate(dir + "/theme/_clean/", pref + "theme/_clean/")
	if err != nil {
		log.Fatal(err)
	}

	err = minimum.walkTemplate(dir + "/theme/_minimum/", pref + "theme/_minimum/")
	if err != nil {
		log.Fatal(err)
	}

	err = hexagonal.walkTemplate(dir + "/theme/_hexagonal/", pref + "theme/_hexagonal/")
	if err != nil {
		log.Fatal(err)
	}

	w, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	archivedOnion := archiveToString(onion)
	archivedMVC := archiveToString(mvc)
	archivedMinimum := archiveToString(minimum)
	archivedClean := archiveToString(clean)
	archivedHexagonal := archiveToString(hexagonal)
	archivedLayout := archiveToString(layout)


	if archivedOnion != "" || archivedMVC != "" {
		_, _ = fmt.Fprintln(w, "// DO NOT EDIT.")
		_, _ = fmt.Fprintln(w, "")
		_, _ = fmt.Fprintln(w, "package tmpl")
		_, _ = fmt.Fprintln(w)
		_, _ = fmt.Fprintln(w, `import "text/template"`)
		_, _ = fmt.Fprintln(w)
		_, _ = fmt.Fprintf(w, "var layout = %q\n", archivedLayout)
		_, _ = fmt.Fprintf(w, "var OnionTmpl = template.Must(template.New"+
			"(\"tmpl\").Delims(`@@`, `@@`).Parse(layout+%q))\n", archivedOnion)
		_, _ = fmt.Fprintf(w, "var MVCTmpl = template.Must(template.New"+
			"(\"tmpl\").Delims(`@@`, `@@`).Parse(layout+%q))\n", archivedMVC)
		_, _ = fmt.Fprintf(w, "var MinimumTmpl = template.Must(template.New"+
			"(\"tmpl\").Delims(`@@`, `@@`).Parse(layout+%q))\n", archivedMinimum)
		_, _ = fmt.Fprintf(w, "var CleanTmpl = template.Must(template.New"+
			"(\"tmpl\").Delims(`@@`, `@@`).Parse(layout+%q))\n", archivedClean)
		_, _ = fmt.Fprintf(w, "var HexagonalTmpl = template.Must(template.New"+
			"(\"tmpl\").Delims(`@@`, `@@`).Parse(layout+%q))\n", archivedHexagonal)
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
