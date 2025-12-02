package textutil

import "testing"

func TestRepeatString_Basic(t *testing.T) {
	got := RepeatString("go", 3)
	want := "gogogo"

	if got != want {
		t.Fatalf("RepeatString(\"go\", 3) = %q; want %q", got, want)
	}
}

func TestRepeatString_Zero(t *testing.T) {
	got := RepeatString("go", 0)
	want := ""

	if got != want {
		t.Fatalf("RepeatString(\"go\", 0) = %q; want %q", got, want)
	}
}

func TestRespeatString_Chinese(t *testing.T) {
	got := RepeatString("哈哈哈哈哈哈哈", 2)
	want := "哈哈哈哈哈哈哈哈哈哈哈哈哈哈"

	if got != want {
		t.Fatalf("RepeatString(\"哈哈哈哈哈哈\", 2) = %q; want %q", got, want)
	}
}
