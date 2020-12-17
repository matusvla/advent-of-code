package convolution

type Convolver3D struct {
	mask         [3][3][3]int
	mappingFuncs []func(int) int
}

func New(mask [3][3][3]int, mf ...func(int) int) Convolver3D {
	return Convolver3D{
		mask:         mask,
		mappingFuncs: mf,
	}
}

func convolute(ini, inj, ink int, original [][][]int, mask [3][3][3]int) int {
	var result int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				result += original[ini-1+i][inj-1+j][ink-1+k] * mask[i][j][k]
			}
		}
	}
	return result
}

func (c Convolver3D) ConvolutionExtendBoundsWithMod(in [][][]int) [][][]int {

	zSize := len(in)
	ySize := len(in[0])
	xSize := len(in[0][0])
	// input allocation
	input := make([][][]int, zSize+4)
	for i := range input {
		input[i] = make([][]int, ySize+4)
		for j := range input[i] {
			input[i][j] = make([]int, xSize+4)
		}
	}
	for i, plane := range in {
		for j, row := range plane {
			for k, val := range row {
				input[i+2][j+2][k+2] = val
			}
		}
	}

	// output allocation
	output := make([][][]int, zSize+2)
	for i := range output {
		output[i] = make([][]int, ySize+2)
		for j := range output[i] {
			output[i][j] = make([]int, xSize+2)
		}
	}

	// convolution
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			for k := 1; k < len(input[i][j])-1; k++ {
				convResult := convolute(i, j, k, input, c.mask)
				output[i-1][j-1][k-1] = c.mappingFuncs[input[i][j][k]](convResult) // todo this can panic
			}
		}
	}
	return output
}
