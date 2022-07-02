// https://go-tour-jp.appspot.com/methods/19

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// 戻り値の error は以下のインタフェース。Error() の実装が必要
// type error interface {
//    Error() string
//}
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

// nil でない場合は Error メソッドを呼び出す。
func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
