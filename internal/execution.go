package internal

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/gari8/sprinter/tools/tmpl"
	"go/build"
	"go/format"
	"golang.org/x/tools/txtar"
	"gopkg.in/AlecAivazis/survey.v1"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	Psql  DBType = "PostgresQL"
	Mysql DBType = "MySQL"
)

const (
	Onion   ArcType = "Onion"
	MVC     ArcType = "MVC"
	Minimum ArcType = "Minimum"
	Clean   ArcType = "Clean"
)

func Exec () {
	var sp Sprinter
	flag.Parse()
	switch flag.Arg(0) {
	case "new":
		sp.isCreateMode = true
	case "help":
		sp.isHelpMode = true
	}


	if sp.isHelpMode {
		PrintAny(PYellow, helpMessage)
		return
	}

	if sp.isCreateMode {
		sp.conversation()
		if err := sp.Run(); err != nil {
			log.Fatal(err)
		}

		PrintAny(PGreen,`
			
			...complete

		`)
		PrintAny(PMagenta,"please enter following command")
		PrintAny(PMagenta,"$ cd "+sp.ImportPath)
		PrintAny(PMagenta,"$ docker-compose up --build")
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

	s.ImportPath = answers.Name

	switch answers.DBN {
	case string(Psql):
		s.DataBase = Psql
	case string(Mysql):
		s.DataBase = Mysql
	default:
		s.DataBase = Psql
	}

	switch answers.Arc {
	case string(Onion):
		s.Mode = Onion
	case string(MVC):
		s.Mode = MVC
	case string(Minimum):
		s.Mode = Minimum
	case string(Clean):
		s.Mode = Clean
	default:
		s.Mode = Minimum
	}

	PrintAny(PGreen,"ok")
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
		sym.DataBase = Psql
	case Mysql:
		sym.DataBase = Mysql
	default:
		PrintAny(PRed,"Enter either 'postgres' or 'mysql' for the database")
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	sym.ImportPath = s.importPath(cwd)

	if sym.ImportPath == "" {
		return fmt.Errorf("%s: package name is not found", s.ImportPath)
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
		if err := tmpl.OnionTmpl.Execute(&buf, sym); err != nil {
			return err
		}
	case MVC:
		if err := tmpl.MVCTmpl.Execute(&buf, sym); err != nil {
			return err
		}
	case Minimum:
		if err := tmpl.MinimumTmpl.Execute(&buf, sym); err != nil {
			return err
		}
	case Clean:
		if err := tmpl.CleanTmpl.Execute(&buf, sym); err != nil {
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

	PrintAny(PGreen, "create: "+fp)

	return nil
}

var helpMessage = `
please enter your console
sprinter -n
then you may play nice time

-help or -h help command
-new or -n create command
`