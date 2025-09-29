package main

import (
	"fmt"
	"github.com/yoshioka0101/sample_go/b"
)

func init() {
	fmt.Println("main package の init 関数が実行されました")
}

func main() {
	fmt.Println("main 関数が開始されました")
	b.CallCFunction()
	fmt.Println("main 関数が終了しました")
}