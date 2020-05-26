package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/iancoleman/strcase"
	"strings"
)

func main() {
	// クリップボードから読み込み
	data, _ := clipboard.ReadAll()
	result := ""
	for _, s := range strings.Split(data, "\n") {
		// スネークケースに変換
		replaced := strcase.ToSnake(s)
		result += replaced + "\n"
	}
	// クリップボードに書き込み
	err := clipboard.WriteAll(result)
	if err != nil {
		fmt.Println(err)
	}
}
