
type indexedStack struct {
contents string
top      int
}

func newIndexedStack(items string) *indexedStack {
return &indexedStack{
contents: items,
top:      len(items) - 1,
}
}

func (is* indexedStack) push(s string) {
is.contents += s
is.top += len(s)
}

func (is *indexedStack) pop(count int) string {
is.top -= count
return is.contents[top]
}