package main

import (
	"fmt"
	//"os"
)

var cmdInit = &Command{
	Run:       runInit,
	UsageLine: "init",
	Short:     "initialize password database",
	Long:      "initialize password database",
}

func runInit(cmd *Command, args []string) {
	// 空の配列jsonをつくる
	//json := "[]"
	fmt.Print("Enter master password: ")
	scanner := Scanner{
		Ask:        "Enter masster password: ",
		Validation: true,
	}
	passwd := scanner.Input()
	fmt.Println(passwd)
	// パスワード入力
	// パスワード入力
	// 暗号化
	// 暗号化した文字列をファイルに書き込む
	// index.jsonを作る
}
