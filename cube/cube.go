package cube

type face [][]int
type cubie []int

// Cubic : contains the slice of face values for each face creating a cube
type Cubic []face

var faceSchema face
var cubieRows cubie

/*
faces as per scheme on wiki--
	red			blue		orange		green		yellow		white
	front 		right		back		left		btm			top
	[0, 0, 0]	[1, 1, 1]	[2, 2, 2] 	[3, 3, 3]	[4, 4, 4]	[5, 5, 5]
	[0, 0, 0]	[1, 1, 1]	[2, 2, 2]	[3, 3, 3]	[4, 4, 4]	[5, 5, 5]
	[0, 0, 0]	[1, 1, 1]	[2, 2, 2]	[3, 3, 3]	[4, 4, 4]	[5, 5, 5]
*/

// GenerateCube : consist of 6 faces
func GenerateCube() Cubic {
	rows := 3
	cols := 3
	faceValues := [6]int{0, 1, 2, 3, 4, 5}

	cube := make(Cubic, 0, 6)

	for i := range faceValues {
		face := buildFace(rows, cols, faceValues[i])
		cube = append(cube, face)

	}
	return cube
}

func buildFace(r, c, faceVal int) face {
	faceSchema = make(face, r)
	cubieRows = make(cubie, r*c) // creates a cubie with len of 9
	// modify cubies to make slices of 9 for 0-5
	for i := 0; i < len(cubieRows); i++ {
		cubieRows[i] = faceVal
	}
	// splits cubieRows into slice of 3 slices
	for j := range faceSchema {
		faceSchema[j] = cubieRows[j*c : (j+1)*c]
	}
	return faceSchema
}
