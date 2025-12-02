package textutil

// RepeatString 将字符串 s 重复 n 次
// 例如： RepeatSrting("go,3) => "gogogo"
func RepeatString(s string, n int) string {
	result := ""

	for i := 0; i < n; i++ {
		result += s
	}

	return result
}
