package stack

import "testing"

func TestIndexedStack(t *testing.T) {
	is := New("hello")
	is.Push("world")
	popped := is.Pop(6)
	if popped != "dlrowo" {
		t.Errorf("expected dlrowo, got %s", popped)
	}
	is.Push("o")
	popped = is.Pop(5)
	if popped != "olleh" {
		t.Errorf("expected olleh, got %s", popped)
	}
}
