package forest

type Forest struct {
	trees [][]int
}

func New() *Forest {
	return &Forest{}
}

func (fst *Forest) AddTreeLine(treeLineStr string) {
	var treeLine []int
	for _, c := range treeLineStr {
		treeLine = append(treeLine, int(c-'0'))
	}
	fst.trees = append(fst.trees, treeLine)
}

func (fst *Forest) CountVisibleTrees() int {
	f := fst.trees
	visible := make([][]bool, len(f))
	for i := range visible {
		visible[i] = make([]bool, len(f[i]))
	}
	for i := 0; i < len(f); i++ {
		highestTree := -1
		for j := 0; j < len(f[i]); j++ {
			treeHeight := f[i][j]
			if treeHeight > highestTree {
				highestTree = treeHeight
				visible[i][j] = true
			}
		}
	}
	// same thing, but right to left
	for i := len(f) - 1; i >= 0; i-- {
		highestTree := -1
		for j := len(f[i]) - 1; j >= 0; j-- {
			treeHeight := f[i][j]
			if treeHeight > highestTree {
				highestTree = treeHeight
				visible[i][j] = true
			}
		}
	}
	// same thing, but top to bottom
	for j := 0; j < len(f[0]); j++ {
		highestTree := -1
		for i := 0; i < len(f); i++ {
			treeHeight := f[i][j]
			if treeHeight > highestTree {
				highestTree = treeHeight
				visible[i][j] = true
			}
		}
	}

	// same thing, but bottom to top
	for j := len(f[0]) - 1; j >= 0; j-- {
		highestTree := -1
		for i := len(f) - 1; i >= 0; i-- {
			treeHeight := f[i][j]
			if treeHeight > highestTree {
				highestTree = treeHeight
				visible[i][j] = true
			}
		}
	}

	// count visible trees
	count := 0
	for i := 0; i < len(f); i++ {
		for j := 0; j < len(f[i]); j++ {
			if visible[i][j] {
				count++
			}
		}
	}

	return count
}

func (fst *Forest) FindBestScenicScore() int {
	f := fst.trees
	scenicScore := make([][]int, len(f))
	for i := range scenicScore {
		scenicScore[i] = make([]int, len(f[i]))
		for j := 0; j < len(scenicScore[i]); j++ {
			scenicScore[i][j] = 1
		}
	}

	for i := 0; i < len(f); i++ {
		for j := 0; j < len(f[i]); j++ {
			treeHeight := f[i][j]
			Lsc, Rsc, Usc, Dsc := 0, 0, 0, 0
			for k := j - 1; k >= 0; k-- {
				Lsc++
				if f[i][k] >= treeHeight {
					break
				}
			}
			for k := j + 1; k < len(f[i]); k++ {
				Rsc++
				if f[i][k] >= treeHeight {
					break
				}
			}
			for k := i - 1; k >= 0; k-- {
				Usc++
				if f[k][j] >= treeHeight {
					break
				}
			}
			for k := i + 1; k < len(f); k++ {
				Dsc++
				if f[k][j] >= treeHeight {
					break
				}
			}
			scenicScore[i][j] = Lsc * Rsc * Usc * Dsc
		}
	}

	max := 0
	for i := 0; i < len(f); i++ {
		for j := 0; j < len(f[i]); j++ {
			if scenicScore[i][j] > max {
				max = scenicScore[i][j]
			}
		}
	}
	return max
}
