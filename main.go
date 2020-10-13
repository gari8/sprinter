package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/build"
	"go/format"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/tools/txtar"
)

//go:generate go run tools/txtar/main.go -strip "_template/" _template template.go

type (
	Sprinter struct {
		ImportPath string
		DataBase string
		ExeName string
		Dir string
		Args []string
	}
	Symbol struct {
		Pkg string
		ImportPath string
		DataBase int
		GoVer string
	}
)

const (
	Psql = iota
	Mysql
)

func main() {
	var sprinter Sprinter
	flag.StringVar(&sprinter.ImportPath, "path", "", "import path")
	flag.StringVar(&sprinter.DataBase, "db", "", "which database default postgres")
	flag.Parse()
	sprinter.ExeName = os.Args[0]
	sprinter.Args = flag.Args()

	if err := sprinter.Run(); err != nil {
		log.Fatal(err)
	}
}

func (s *Sprinter) Run() error {
	sym := &Symbol{
		GoVer:   runtime.Version()[2:],
	}

	if s.ImportPath != "" {
		s.Dir = s.ImportPath
		sym.Pkg = path.Base(s.ImportPath)
	} else {
		return errors.New("package name is not found")
	}

	if s.DataBase != "" {
		switch s.DataBase {
		case "postgres":
			sym.DataBase = Psql
		case "mysql":
			sym.DataBase = Mysql
		default:
			log.Fatal("Enter either 'postgres' or 'mysql' for the database")
		}
	} else {
		fmt.Println("The database basically uses postgres")
		fmt.Println("If you want to use mysql, select mysql with '-db' option")
		sym.DataBase = Psql
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	sym.ImportPath = s.importPath(cwd)

	if sym.ImportPath == "" {
		return fmt.Errorf("%s: package name is not found", s.ExeName)
	}

	if err := s.createAll(sym); err != nil {
		return err
	}


	return nil
}


func (s *Sprinter) importPath(cwd string) string {

	if s.ImportPath != "" {
		return s.ImportPath
	}

	for _, gopath := range filepath.SplitList(build.Default.GOPATH) {
		if gopath == "" {
			continue
		}

		src := filepath.Join(gopath, "src")
		if strings.HasPrefix(cwd, src) {
			rel, err := filepath.Rel(src, cwd)
			if err != nil {
				return ""
			}
			return path.Join(filepath.ToSlash(rel), filepath.ToSlash(s.Dir))
		}
	}

	return ""
}

func (s *Sprinter) createAll(sym *Symbol) error {

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, sym); err != nil {
		return err
	}

	ar := txtar.Parse(buf.Bytes())
	for _, f := range ar.Files {
		if err := s.createFile(f); err != nil {
			return err
		}
	}

	return nil
}

func (s *Sprinter) createFile(f txtar.File) (rerr error) {
	if len(bytes.TrimSpace(f.Data)) == 0 {
		return nil
	}

	fp := filepath.Join(s.Dir, filepath.FromSlash(f.Name))


	if err := os.MkdirAll(filepath.Dir(fp), 0777); err != nil {
		return err
	}

	w, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer func() {
		if err := w.Close(); err != nil && rerr == nil {
			rerr = err
		}
	}()

	// format a go file
	data := f.Data
	if filepath.Ext(fp) == ".go" {
		data, err = format.Source(data)
		if err != nil {
			return err
		}
	}

	r := bytes.NewReader(data)
	if _, err := io.Copy(w, r); err != nil {
		return err
	}

	fmt.Println("create", fp)

	return nil
}
