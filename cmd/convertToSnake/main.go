package main

import (
	"github.com/i-chi-li/convertStrCases"
	"github.com/iancoleman/strcase"
)

func main() {
	// スネークケースに変換
	convertStrCases.ConvertTo(strcase.ToSnake)
}
