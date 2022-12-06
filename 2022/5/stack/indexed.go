package stack

type IndexedStack struct {
	contents    string
	top         int
	isReverting bool
}

func New(items string, isReverting bool) *IndexedStack {
	return &IndexedStack{
		contents:    items,
		top:         len(items) - 1,
		isReverting: isReverting,
	}
}

func (is *IndexedStack) Push(s string) {
	is.contents = is.contents[:is.top+1] + s
	is.top += len(s)
}

func (is *IndexedStack) Pop(count int) string {
	is.top -= count
	toPop := is.contents[is.top+1 : is.top+count+1]
	if is.isReverting {
		return reverse(toPop)
	}
	return toPop
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func (is *IndexedStack) Top() string {
	if is == nil || is.top == -1 {
		return ""
	}
	return string(is.contents[is.top])
}

func (is *IndexedStack) String() string {
	return is.contents
}
