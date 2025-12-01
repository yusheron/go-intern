package textutil

import "strings"

// WordCount 统计一行英文文本中每个单词出现的次数
// 简单版本，只考虑空格分隔的单词，忽略标点符号和大小写
func WordCount(line string) map[string]int {
	result := make(map[string]int)

	// strings.Fields 会根据空白字符分割字符串
	fields := strings.Fields(line)
	for _, word := range fields {
		result[word]++
	}

	return result
}
