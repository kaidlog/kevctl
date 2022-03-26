package cmd

import (
	"log"
	"strings"

	"github.com/go-programming-tour-book/tour/internal/word"
	"github.com/spf13/cobra"
)

const (
	ModeUpper                      = iota + 1 // 全部轉大寫
	ModeLower                                 // 全部轉小寫
	ModeUnderscoreToUpperCamelCase            // 下劃線轉大寫駝峰
	ModeUnderscoreToLowerCamelCase            // 下線線轉小寫駝峰
	ModeCamelCaseToUnderscore                 // 駝峰轉下劃線
)

var str string
var mode int8
var desc = strings.Join([]string{
	"該子命令支持各種單詞格式轉換，模式如下：",
	"1：全部轉大寫",
	"2：全部轉小寫",
	"3：下劃線轉大寫駝峰",
	"4：下劃線轉小寫駝峰",
	"5：駝峰轉下劃線",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "單詞格式轉換",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暫不支持該轉換模式，請執行 help word 查看幫助文檔")
		}

		log.Printf("輸出結果: %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "請輸入單詞內容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "請輸入單詞轉換的模式")
}
