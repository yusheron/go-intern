package textutil

import "testing"

// 测试一个最简单的句子
func TestWordCount_Simple(t *testing.T) {
	got := WordCount("hello world hello")

	if got["hello"] != 2 {
		t.Fatalf("expected hello=2, got %d", got["hello"])
	}
	if got["world"] != 1 {
		t.Fatalf("expected world=1, got %d", got["world"])
	}
}

// 测试空字符串
func TestWordCount_Empty(t *testing.T) {
	got := WordCount("")

	if len(got) != 0 {
		t.Fatalf("expected empty map, got %#v", got)
	}
}

// 测试多个空格和制表符
func TestWordCount_MultiSpaces(t *testing.T) {
	got := WordCount("foo   bar\tfoo")

	if got["foo"] != 2 {
		t.Fatalf("expected foo=2, got %d", got["foo"])
	}
	if got["bar"] != 1 {
		t.Fatalf("expected bar=1, got %d", got["bar"])
	}
}
