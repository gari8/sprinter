package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/build"
	"go/format"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/tools/txtar"
)

//go:generate go run tools/txtar/main.go -strip "_template/" _template template.go

var archive txtar.Archive

type (
	Sprinter struct {
		OverWrite bool
		ImportPath string
		ExeName string
		Dir string
		Args []string
	}
	Symbol struct {
		Pkg string
		ImportPath string
		GoVer string
	}
)

func main() {
	var sprinter Sprinter
	flag.BoolVar(&sprinter.OverWrite, "overwrite", false, "overwrite everything")
	flag.StringVar(&sprinter.ImportPath, "path", "", "import path")
	flag.Parse()
	sprinter.ExeName = os.Args[0]
	sprinter.Args = flag.Args()

	if err := sprinter.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func (s *Sprinter) Run() error {
	sym := &Symbol{
		GoVer:   runtime.Version()[2:],
	}

	if len(s.Args) < 1 {
		if s.ImportPath != "" {
			s.Dir = s.ImportPath
			sym.Pkg = path.Base(s.ImportPath)
		} else {
			return errors.New("package can't be made here")
		}
	} else {
		s.Dir = s.Args[0]
		sym.Pkg = path.Base(s.Args[0])
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	sym.ImportPath = s.importPath(cwd)

	if sym.ImportPath == "" {
		const format = "%s must be executed in GOPATH or -path option must be specified"
		return fmt.Errorf(format, s.ExeName)
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
	//if err := tmpl.Execute(&buf, sym); err != nil {
	//	return err
	//}

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

	path := filepath.Join(s.Dir, filepath.FromSlash(f.Name))


	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		return err
	}

	w, err := os.Create(path)
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
	if filepath.Ext(path) == ".go" {
		data, err = format.Source(data)
		if err != nil {
			return err
		}
	}

	r := bytes.NewReader(data)
	if _, err := io.Copy(w, r); err != nil {
		return err
	}

	fmt.Println("create", path)

	return nil
}
