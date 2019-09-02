package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func (it *Item) displayPath() {

	if it.parent != nil {
		it.parent.displayPath()
		fmt.Printf("%v|\n", strings.Repeat(" ", matrixDisplayWidth/2))
		fmt.Printf("%vv\n", strings.Repeat(" ", matrixDisplayWidth/2))
	}

	it.m.display()
}

func (m Matrix) display() {

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
	fmt.Println(buff.String())
}

func (m Matrix) string() string {

	var buff bytes.Buffer
	for k, row := range m {
		if k != 0 {
			buff.WriteByte('\n')
		}

		for p, num := range row {
			if p != 0 {
				buff.WriteByte(' ')
			}

			buff.WriteString(strconv.Itoa(num))
		}
	}

	return buff.String()
}
