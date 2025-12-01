package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yusheron/go-intern/pkg/textutil"
)

func main() {
	// 创建一个从标准输入读取的 Scanner
	scanner := bufio.NewScanner(os.Stdin)

	// 读取一行，如果没有读到（比如直接 Ctrl+Z 回车），就退出
	if !scanner.Scan() {
		return
	}

	// 拿到这一行的文本
	line := scanner.Text()

	// 调用我们写好的 ReverseString
	reversed := textutil.ReverseString(line)

	// 打印结果
	fmt.Println(reversed)
}
