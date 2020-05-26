package main

import (
	"github.com/atotto/clipboard"
	"github.com/iancoleman/strcase"
	"strings"
)

func main() {
	// クリップボードから読み込み
	data, _ := clipboard.ReadAll()
	// 変換後文字列
	result := ""
	for i, s := range strings.Split(data, "\n") {
		if i != 0 {
			// 最初の行以外の場合
			result += "\n"
		}
		// スネークケースに変換
		result += strcase.ToSnake(s)
	}
	// クリップボードに書き込み
	_ = clipboard.WriteAll(result)
}
