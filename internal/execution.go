package internal

import (
	"flag"
	"log"
)

const (
	Psql  DBType = "PostgresQL"
	Mysql DBType = "MySQL"
)

const (
	Onion     ArcType = "Onion"
	MVC       ArcType = "MVC"
	Minimum   ArcType = "Minimum"
	Clean     ArcType = "Clean"
	Hexagonal ArcType = "Hexagonal"
)

func Exec() {
	var sp Sprinter
	flag.Parse()
	switch flag.Arg(0) {
	case "new":
		sp.isCreateMode = true
	case "help":
		sp.isHelpMode = true
	}

	if sp.isHelpMode {
		PrintAny(PYellow, help)
		return
	}

	if sp.isCreateMode {
		sp.conversation()
		if err := sp.Run(); err != nil {
			log.Fatal(err)
		}

		PrintAny(PGreen, `
			
			...complete

		`)
		PrintAny(PCyan, "please enter following command")
		PrintAny(PCyan, "$ cd "+sp.ImportPath)
		PrintAny(PCyan, "$ make init")
		return
	}

	PrintAny(PBlue, guide)
}

var guide = `
please enter your console

sprinter new
- creating new app

sprinter help
- helping you use this tool
`

var help = `
creating new app
$ sprinter new

To know more
https://github.com/gari8/sprinter
`
