package main

import (
	"github.com/i-chi-li/convertStrCases"
	"github.com/iancoleman/strcase"
)

func main() {
	// アッパーキャメルケースに変換
	convertStrCases.ConvertTo(strcase.ToCamel)
}
