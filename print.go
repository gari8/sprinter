package sprinter

import "fmt"

const (
	PRed = iota + 31
	PGreen
	PYellow
	PBlue
	PMagenta
	PCyan
)

func PrintAny(number int, str string) {
	fmt.Printf("\x1b[%dm%s\x1b[0m\n", number, str)
}
