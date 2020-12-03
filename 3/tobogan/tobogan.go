package tobogan

const tree = '#'

func CheckTree(line string, lineNo int, step float32) bool {
	gotoPlace := step * float32(lineNo)
	return gotoPlace == float32(int(gotoPlace)) && line[int(gotoPlace)%len(line)] == tree
}
