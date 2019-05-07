package cube

import "fmt"

type face [][]int
type row []int

// Cubic : contains the slice of face values for each face creating a cube
type Cubic []face

/*
faces as per scheme on wiki--
	red			blue		orange		green		yellow		white
	front 		right		back		left		down		up
	[0, 0, 0]	[1, 1, 1]	[2, 2, 2] 	[3, 3, 3]	[4, 4, 4]	[5, 5, 5]
	[0, 0, 0]	[1, 1, 1]	[2, 2, 2]	[3, 3, 3]	[4, 4, 4]	[5, 5, 5]
	[0, 0, 0]	[1, 1, 1]	[2, 2, 2]	[3, 3, 3]	[4, 4, 4]	[5, 5, 5]
*/

// GenerateCube : Returns a new cube that is solved
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
	faceSchema := make(face, r) // creates empty face schema
	cubieFace := make(row, r*c) // creates all cubies in face; len = 9
	for i := range cubieFace {
		cubieFace[i] = faceVal // fills rows with correct face vals
	}

	// splits cubieFace into slice of 3 slices and adds them to
	// faceSchema creating a generated face
	for j := range faceSchema {
		faceSchema[j] = cubieFace[j*c : (j+1)*c]
	}
	return faceSchema
}

//PrintCube : Formats to Stdout Cube Visualization
func (c Cubic) PrintCube() {
	fmt.Printf("Here Is Your Cube: \n")

	pos := []string{"F", "R", "B", "L", "D", "U"}

	for i, v := range pos {
		fmt.Printf("%v: %v\n", v, c[i])
	}
}
