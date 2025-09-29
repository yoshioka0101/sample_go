package c

import "fmt"

func init() {
	fmt.Println("package c の init 関数が実行されました")
}

func CallFromC() {
	fmt.Println("package c の関数が呼び出されました")
}