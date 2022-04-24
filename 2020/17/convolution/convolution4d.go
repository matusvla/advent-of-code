package convolution

type Convolver4D struct {
	mask         [3][3][3][3]int
	mappingFuncs []func(int) int
}

func New4D(mask [3][3][3][3]int, mf ...func(int) int) Convolver4D {
	return Convolver4D{
		mask:         mask,
		mappingFuncs: mf,
	}
}

func convolute4d(ini, inj, ink, inl int, original [][][][]int, mask [3][3][3][3]int) int {
	var result int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					result += original[ini-1+i][inj-1+j][ink-1+k][inl-1+l] * mask[i][j][k][l]
				}
			}
		}
	}
	return result
}

func (c Convolver4D) ConvolutionExtendBoundsWithMod(in [][][][]int) [][][][]int {

	zSize := len(in)
	wSize := len(in[0])
	ySize := len(in[0][0])
	xSize := len(in[0][0][0])
	// input allocation
	input := make([][][][]int, zSize+4)
	for i := range input {
		input[i] = make([][][]int, wSize+4)
		for j := range input[i] {
			input[i][j] = make([][]int, ySize+4)
			for k := range input[i][j] {
				input[i][j][k] = make([]int, xSize+4)
			}
		}
	}
	for h, hyperplane := range in {
		for i, plane := range hyperplane {
			for j, row := range plane {
				for k, val := range row {
					input[h+2][i+2][j+2][k+2] = val
				}
			}
		}
	}

	// output allocation
	output := make([][][][]int, zSize+2)
	for i := range output {
		output[i] = make([][][]int, wSize+2)
		for j := range output[i] {
			output[i][j] = make([][]int, ySize+2)
			for k := range output[i][j] {
				output[i][j][k] = make([]int, xSize+2)
			}
		}
	}

	// convolution
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			for k := 1; k < len(input[i][j])-1; k++ {
				for l := 1; l < len(input[i][j])-1; l++ {
					convResult := convolute4d(i, j, k, l, input, c.mask)
					output[i-1][j-1][k-1][l-1] = c.mappingFuncs[input[i][j][k][l]](convResult) // todo this can panic
				}
			}
		}
	}
	return output
}
