package helper

import (
	"github.com/davecgh/go-spew/spew"
	"log"
)

func DumpLog(res interface{}) {
	log.Println(spew.Sdump(res))
}

// credits: https://stackoverflow.com/questions/28008566/how-to-compute-the-offset-from-column-and-line-number-go
func FindOffset(fileText string, line, column int) int {
	currentCol := 1
	currentLine := 1

	for offset, ch := range fileText {
		if currentLine == line && currentCol == column {
			return offset
		}

		if ch == '\n' {
			currentLine++
			currentCol = 1
		} else {
			currentCol++
		}

	}
	return -1
}
