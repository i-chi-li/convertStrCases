package main

import (
	"github.com/i-chi-li/convertStrCases"
	"github.com/iancoleman/strcase"
)

func main() {
	// ケバブケースに変換
	convertStrCases.ConvertTo(strcase.ToKebab)
}
