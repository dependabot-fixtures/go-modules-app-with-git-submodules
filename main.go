package main

import (
	examplelib "github.com/dependabot-fixtures/go-modules-lib"
	"github.com/fatih/color"
)

func main() {
	color.Cyan(examplelib.Greet())
}
