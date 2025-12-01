package textutil

import "testing"

// 测试英文字符串反转
func TestReverseString_English(t *testing.T) {
	got := ReverseString("hello")
	want := "olleh"

	if got != want {
		t.Fatalf("ReverseString(\"hello\") = %q; want %q", got, want)
	}
}

// 测试包含中文的字符串反转
func TestReverseString_Chinese(t *testing.T) {
	got := ReverseString("你好Go")
	want := "oG好你"

	if got != want {
		t.Fatalf("ReverseString(\"你好Go\") = %q; want %q", got, want)
	}
}

// 测试空字符串
func TestReverseString_Empty(t *testing.T) {
	got := ReverseString("")
	want := ""

	if got != want {
		t.Fatalf("ReverseString(\"\") = %q; want %q", got, want)
	}
}
