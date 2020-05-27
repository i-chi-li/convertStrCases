package convertStrCases

import (
	"fmt"
	"github.com/atotto/clipboard"
	"regexp"
)

// 単語部分のみを変換対象とする
var re = regexp.MustCompile(`[^ \n\r,"'()\[\]{}]+`)

func ConvertTo(repl func(string) string) {
	// クリップボードから読み込み
	data, _ := clipboard.ReadAll()
	// 置換用関数
	f := func(s string) string {
		fmt.Printf("[%s]\n", s)
		// ケバブケースに変換
		return repl(s)
	}
	result := re.ReplaceAllStringFunc(data, f)
	fmt.Printf("[%s]", result)
	// クリップボードに書き込み
	_ = clipboard.WriteAll(result)
}
