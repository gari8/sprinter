package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/build"
	"go/format"
	"gopkg.in/AlecAivazis/survey.v1"
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
		DataBase int
		ExeName string
		Dir string
		isCreateMode bool
		isHelpMode bool
		Mode int
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

const (
	Onion = iota + 1
	MVC
)

var qs = []*survey.Question{
	{
		Name: "name",
		Prompt: &survey.Input{
			Message: "[1] Please enter the title of your application",
		},
		Validate: survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "dbname",
		Prompt: &survey.Select{
			Message: "[2] Please select the database",
			Options: []string{"Postgres", "Mysql"},
			Default: "Postgres",
		},
	},
	{
		Name: "architecture",
		Prompt: &survey.Select{
			Message: "[3] Please select the architecture",
			Options: []string{"Onion", "MVC"},
		},
	},
}

func main() {
	var sprinter Sprinter
	flag.BoolVar(&sprinter.isCreateMode, "new", false, "create mode")
	flag.BoolVar(&sprinter.isHelpMode, "help", false, "help mode")
	flag.BoolVar(&sprinter.isCreateMode, "n", false, "create mode")
	flag.BoolVar(&sprinter.isHelpMode, "h", false, "help mode")
	flag.Parse()
	sprinter.ExeName = os.Args[0]

	if sprinter.isHelpMode {
		fmt.Print(helpMessage)
		return
	}

	if sprinter.isCreateMode {
		sprinter.conversation()
		if err := sprinter.Run(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(`
			
			...complete
		`)
		return
	}

	fmt.Println("\n-h or -help")
}

func (s *Sprinter) conversation() {
	answers := struct {
		Name          string `survey:"name"`
		DBN           string `survey:"dbname"`
		Arc           string `survey:"architecture"`
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		log.Fatalf("[error] %s  ...stopping", err)
	}

	if answers.Name == "" {
		return
	}

	s.ImportPath = strings.ToLower(answers.Name)

	switch answers.DBN {
	case "Postgres":
		s.DataBase = Psql
	case "Mysql":
		s.DataBase = Mysql
	default:

	}

	switch answers.Arc {
	case "Onion":
		s.Mode = Onion
	case "MVC":
		s.Mode = MVC
	default:

	}

	fmt.Println("ok")
	fmt.Println("")
	fmt.Println("")
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

	switch s.DataBase {
	case Psql:
		sym.DataBase = s.DataBase
	case Mysql:
		sym.DataBase = s.DataBase
	default:
		log.Fatal("Enter either 'postgres' or 'mysql' for the database")
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
	switch s.Mode {
	case Onion:
		if err := tmplOnion.Execute(&buf, sym); err != nil {
			return err
		}
	case MVC:
		if err := tmplMVC.Execute(&buf, sym); err != nil {
			return err
		}
	default:
		return errors.New("invalid command")
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

var helpMessage = `
-help or -h help command
-new or -n create command
`
