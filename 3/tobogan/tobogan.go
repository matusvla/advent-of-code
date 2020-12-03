package tobogan

const (
	tree = '#'
)

func CheckTree(line string, lineNo int, step float32) bool {
	gotoPlace := step * float32(lineNo)
	if gotoPlace != float32(int(gotoPlace)) {
		return false
	}
	return line[int(gotoPlace)%len(line)] == tree
}
