package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yusheron/go-intern/pkg/textutil"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		//没有输入就直接退出
		return
	}

	line := scanner.Text()

	freq := textutil.WordCount(line)
	for word, count := range freq {
		fmt.Println(word, count)
	}
}
