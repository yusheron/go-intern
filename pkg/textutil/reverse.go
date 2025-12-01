package textutil

// ReverseString 反转字符串，使用 rune 切片以支持中文等多字节字符。
func ReverseString(s string) string {
	// 将字符串转成 rune 切片，避免中文字符被拆开
	runes := []rune(s)
	n := len(runes)

	// 头尾交换
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		runes[i], runes[j] = runes[j], runes[i]
	}

	// 再转回字符串
	return string(runes)
}
