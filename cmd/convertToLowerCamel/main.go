package main

import (
	"github.com/i-chi-li/convertStrCases"
	"github.com/iancoleman/strcase"
)

func main() {
	// ローワーキャメルケースに変換
	convertStrCases.ConvertTo(strcase.ToLowerCamel)
}
