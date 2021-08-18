package internal

import (
	"gopkg.in/AlecAivazis/survey.v1"
)

type (
	DBType string
	ArcType string
	Sprinter struct {
		ImportPath   string
		DataBase     DBType
		Dir          string
		isCreateMode bool
		isHelpMode   bool
		Mode         ArcType
	}
	Symbol struct {
		Pkg        string
		ImportPath string
		DataBase   DBType
		GoVer      string
	}
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
			Options: []string{string(Psql), string(Mysql)},
			Default: string(Psql),
		},
	},
	{
		Name: "architecture",
		Prompt: &survey.Select{
			Message: "[3] Please select the architecture",
			Options: []string{string(Clean), string(Hexagonal), string(Minimum), string(MVC), string(Onion)},
		},
	},
}
