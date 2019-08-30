package main

import (
	"bytes"
	"fmt"
	"strings"
)

func (it *Item) displayPath() {

	if it.parent != nil {
		it.parent.displayPath()
		fmt.Printf("%v|\n", strings.Repeat(" ", matrixDisplayWidth/ 2))
		fmt.Printf("%vv\n", strings.Repeat(" ", matrixDisplayWidth/ 2))
	}

	display := it.m.display()
	fmt.Println(display.String())
}

func (m Matrix) display() bytes.Buffer {

	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("%v\n", strings.Repeat("-", matrixDisplayWidth)))
	for _, row := range m {
		buff.WriteString("|")
		for k, num := range row {
			if k != 0 {
				buff.WriteString(" ")
			}
			if num == 0 {
				buff.WriteString("\033[1;33m")
			}
			buff.WriteString(fmt.Sprintf("%*v", itemDisplayWidth, num))
			if num == 0 {
				buff.WriteString("\033[0m")
			}
		}

		buff.WriteString("|\n")
	}

	buff.WriteString(strings.Repeat("-", matrixDisplayWidth))
	return buff
}
