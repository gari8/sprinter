package main

import (
	"testing"

	"golang.org/x/tools/txtar"

	"github.com/stretchr/testify/assert"
)

const (
	pkgName  = "sprinter"
	fileName = "main.go"
	fileData = "package main\nfunc main () {\nprintln(\"sprinter\")\n}"
)

func Test_archiveToString(t *testing.T) {
	file := txtar.File{
		Name: fileName,
		Data: []byte(fileData),
	}
	arc := archive{
		Comment: nil,
		Files:   nil,
	}
	arc.archiveToString()
	assert.Equal(t, "", arc.archiveToString())
	arc = archive{
		Comment: []byte(pkgName),
		Files:   nil,
	}
	assert.Equal(t, pkgName+"\n", arc.archiveToString())
	expected := "sprinter\n-- main.go --\npackage main\nfunc main () {\nprintln(\"sprinter\")\n}\n"
	arc = archive{
		Comment: []byte(pkgName),
		Files:   []txtar.File{file},
	}
	assert.Equal(t, expected, arc.archiveToString())
}

func Test_walkTemplate(t *testing.T) {
	testCaseFile := "../../test/testcase/case_1.txt"
	expected := "-- ../../test/testcase/case_1.txt --\n--NO EDIT--Test--\n"
	arc := archive{
		Comment: nil,
		Files:   nil,
	}
	err := arc.walkTemplate(testCaseFile, "")
	assert.NoError(t, err)
	assert.Equal(t, expected, arc.archiveToString())
}
