package console

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

// 打印成功消息, 绿色输出
func Success(msg string) {
	colorOut(msg, "green")
}

// 打印错误消息, 红色输出
func Error(msg string) {
	colorOut(msg, "red")
}

// 打印警告消息, 黄色输出
func Warning(msg string) {
	colorOut(msg, "yellow")
}

// 打印一条错误消息, 并 os.Exit(1)退出
func Exit(msg string) {
	Error(msg)
	os.Exit(1)
}

// 判断 err != nil
func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}

// 内部使用，使用高亮颜色
func colorOut(msg string, color string) {
	fmt.Fprintln(os.Stdout, ansi.Color(msg, color))
}
