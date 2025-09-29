package b

import (
	"fmt"
	"github.com/yoshioka0101/sample_go/c"
)

func init() {
	fmt.Println("package b の init 関数が実行されました")
}

func CallCFunction() {
	c.CallFromC()
}