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
