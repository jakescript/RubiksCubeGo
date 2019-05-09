# _package cube documentation:_

_this package is built off the idea of displaying a 3-Dimensional object using abstractions of basic types_  
***

### Rubik's Cube Breakdown:  
Cubes cosist of six faces (front, back, left, right, top, and bottom) resulting in a 3-dimensional cube    
Each face is made up of a total of 9 cubies (cube faces)    
A face consists of three types of possible cubes:
 * 4 Corner Cubes 
 * 4 Edge Cubes
 * 1 Static Center Cubes  

In order to work with a 3-dimensional object in go I determined the best approach to be breaking down the cube into slices of slices containing three values. Essentially building one face of the cube and to repeat the process 6 times adding each face to a new slice resulting in a final solved cube.  

I simplified the process of dealing with character and string values by substituting for integers. So in the case of my program, following the cube scheme on wiki, these values correspond to the original color values: 

```
	red       blue      orange    green     yellow    white
	front     right     back      left      btm       top
	[0, 0, 0] [1, 1, 1] [2, 2, 2] [3, 3, 3] [4, 4, 4] [5, 5, 5]
	[0, 0, 0] [1, 1, 1] [2, 2, 2] [3, 3, 3] [4, 4, 4] [5, 5, 5]
	[0, 0, 0] [1, 1, 1] [2, 2, 2] [3, 3, 3] [4, 4, 4] [5, 5, 5]
```

I find it easier to refrence the diagram I designed since the program really interperates the slice as one single line.  
As you can see each face has 3 slices of 3 cubes. Each one of the face values is a slice itself, each face then stored in a final slice of slices; the cube. Yes, lots of slice's of slice's but the point of this project was to practice working with Types and manipulate the abstractions provided from Go. Stick with me here.  

Finally now that I have explained some of the design concepts behind the program how I approach the problem let's get into the code and the package itself.  

***  

## Declaring Types:  
Starting off the package there are some type declarations:

```go
package cube

type face [][]int
type row []int

type Cubic []face
```

The reason I provided the background knowledge on the package is to help understand the first few lines. I created the package cube and this is where I do anything and everything involving the cube to export to the main function later. The first type I create is a _**face**_ and like I explained before, I went with the option of interpreting each color value to be an int, also this is why I created the diagram so you can visualize the multidimensional slice that is a _**face**_. In order to be considered of Type _**face**_, the value must be a slice of slices containing integers. Look at the diagram, a slice of slices which are filled with zero - five, simple each qualifying as a _**face**_.    

Moving on there is a type decleration for a _**row**_, which is of type slice of integers. Later in the code we will see how the _**row**_ type interacts with the _**face**_ type and result in a _**Cubic**_ type. A _**row**_ is litearlly each row of integers in each _**face**_. 

The finaly type we have is our _**Cubic**_ type. This is the simplest to understand and yet the most complex in this program. A basic understanding is that the _**Cubic**_ is that it's made up of a slice of _**face**_, so exactly what the diagram shows, a slice of 6 _**face**_ values (6 being the number of sides on the 3-dimensional cube). A deeper understanding being that in order to be a _**Cubic**_ you need to be a slice of _**face**_, and a slice of _**face**_ is actually a slice of slices containing integers, again _multidimensional arrays_, I find it helps to have a piece of paper near by and writing things down. This project has taken up 10+ pages in my note book figuring out logic and efficiency ( probably an excessive amount of over thinking problems and solutions).  

Now that we understand what makes up the _cube_ lets get into how each _**face**_ is built and then generated into a 3-dimensional cube.  

## Generating A Cube:  
Now we will start to build the foundation of the cube by creating the _**face**_. We already know a cube consist of 6 total faces and in order to be a _**face**_ for a 3x3x3 rubik's cube we need 9 total _cubies_ per face. The main part of the program takes place in the first two functions, **GenerateCube()** & **buildFace(r, c, faceVal int)**. Lets start by following the control flow of the program. We begin in our _main.go_ file in the _main function_, and there we can see a call to _c.GenerateCube()_ so we head back into our _cube package_ and look at what's happening from there.  

```go
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
```  

Right off the function signature we can see that once we use this function we are left with some type of _**Cubic**_. Succeeding that we see three declared variables that only remain in the scope of the _GenerateCube()_ function.  

The first and second are quite obvious being the _rows_ and _columns_ needed for each _**face**_ (remember that each face totals 9 cubies). The third is an array of _faceValues_, the reason this is an array and not a slice is because we know exactly how many faces make up a cube (6). For the purpose of memorory efficiency we can use an array to declare an immutable list of what each _faceValue_ will be (indexed at zero).  

The fourth is special, the variable _cube_ is a slice of _**Cubic**_ with a length of zero and a capacity of 6. Which means our variable _cube_ is going to be of type _**Cubic**_ (which is a slice of face, which is a slice of slices of integers). Read that a couple times over slowly and eventually it won't be hard to conceptualize in your head. The reason for using _make()_ is to preallocate the memory for a slice with a capacity of 6, but we need an empty slice to fill with our faces. I recommend reading the documentation regarding _make()_ if you don't understand it and come back because this pattern is repeated.

Finally, now that we have setup everything we will need to create a _**face**_ we can actually begin the process of molding these variables into each _**face**_, the for loop. Lets look at it again and follow whats happening:  
```go
for i := range faceValues {
  face := buildFace(rows, cols, faceValues[i])
  cube = append(cube, face)
  }
return cube
```

It's a very simple solution to using a function dynamically to create multiple face's of the same structure with different values. We use our _faceValues_ array as a counter for how many face's we want to create, and for each _**face**_ value we set a variable to the return value of the function _buildFace(rows, cols, faceValues[i])_ which we will get to in a minute just stay with me. Next we set _cube_ to a new updated slice containing a _**face**_ per _faceValue_. _The reason for storing the append function and not just calling it is because the function returns the new updated slice with what you appended_ (this is another function if you are unclear of I would read the documentation for), lastly returning the _cube_ variable giving the function the value of _**Cubic**_. To prove this you can go into the main.go file and run this with the proper imports and you should see _cube.Cubic_ as the type:  
```go
func main() {
	c := cube.GenerateCube()
	fmt.Printf("%T\n", c) // %T prints type 
}
```

### Building The Face's  
Now we have ran through the loop that actually builds the face's and compiles all of them into a slice of type _**Cubic**_. Earlier in the loop we came across a line of code that said _face := buildFace(rows, cols, faceValues[i])_. This section is about the _buildFace_ function which is defined as followed:  

```go
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
```

Starting off we have a similar function signature, accepts three integers and returns type _**face**_. Hmmm, type _**face**_? That looks familiar. Going back up to our loop you can see we initialize the _**face**_ variable to be the return value of _buildFace_ which is  of the type _**face**_. So hopefully at this point you can kind of see how all the pieces start to come together for building the _cube_ based off slices, types and a **different perception** on creating using go. I recommend if you are a little lost or unsure you go back to the code and read along line by line to yourself and try writing things down to give yourself a visualization and repetition. If you are all clear and ready to continue lets digest the function and break it down completely.  

_buildFace_ is supplied with three parameters in order to operate, _r being rows, c being columns, faceVal being the digit to represent each face supplied by the for loop (faceValues[i])_. The first two lines we use the _make()_ again delare and allocate the memory needed for two slices of different types. Again here you need to just stick with me and I promise it will make sense as we go along.  
```go
faceSchema := make(face, r) // creates empty face schema
cubieFace := make(row, r*c) // creates all cubies in face; len = 9
```
First up we see _faceSchema_, this becomes a new slice of _**face**_ with a length of three. Looking back up at the types we see a _**face**_ is a type of [][]int (slice containing multiple slices containing integers). Okay so we now have a slice with three other slices that contain zero values. Next up we delcare _cubieFace_ to be a slice of type _**row**_ _(which is essentially a slice of integers, row is just more idiomatic)_ this time with a length equal to _r*c_ (rows * columns) resulting in the total amount of cubies per _**face**_. Right now both of these slices declared have different lengths but actually only contain the zero value of their types. We want to make 6 different faces each with different values so we can now initiate the values to each index in our cubieFace. I sense another for loop!  
```go
for i := range cubieFace {
  cubieFace[i] = faceVal // fills rows with correct face vals
}
```  
This is a simple one to follow we take each index of _cubieFace_ which, remember, is not empty but contains the zero value's of their type _(What is a zero value? Variables declared without an explicit initial value are given their zero value)_. So in the loop we are now setting each index to the faceVal we passed in originally (0-5).  

We finally have our last bit of code! The end of this function has one last for loop and then we will recap quickly. 
```go
for j := range faceSchema {
  faceSchema[j] = cubieFace[j*c : (j+1)*c]
}
return faceSchema
```

This one might be a little harder to follow but lets dive into it. Obviously again we are ranging over a slice, this time our _faceSchema_ (which has a length of 3) so we will go through this loop three times. Each time taking one of the three slices in the _faceSchema_ and setting it to a part of our _cubeFace_ slice. So lets visualize this real quick, 

```
faceSchema = [ [], [], [] ] <- we begin with this empty multidimensional slice which looks like our type face...  
cubieFace = [0, 0, 0, 0, 0, 0, 0, 0, 0] <- now you can visualize a 'zero value'  
```
Back into our loop here we can make a visualization as well to follow along using the first iteration as an example:    
```
for j := range faceSchema { <- j becomes 0 being the first index in faceSchema
faceSchema[j] = cubieFace[j*c : (j+1)*c] <- faceSchema[0] = [] = cubieFace[0*3: (0+1)*3]  
[] : cubieFace[0:3] <- breakdown to index 0 up to index 3 of cubieFace slice
[] : [0,0,0]... <- index 0 of faceSchema is now a row of 3 integers.
```
Following iteration 1 we see that we take our first slice in _faceSchema_ and from a high level view replacing it with index 0 up to but not including index 3 _(From a lower level we are not actually 'replacing' rather 'pointing' to a piece of the cubieFace slice in memory instead)_. Essentially this for loop is taking all of the empty slices in _faceSchema_ and _'building'_ a schema for a _**face**_ by filling each slice with three integers. Finally we are left with a _faceSchema_ that looks like the diagram below:  
```
faceSchema = [ [0,0,0],[0,0,0],[0,0,0] ]  
```
If we follow this pattern back in our _GenerateCube()_ we see this is repeated 6 times each with a _faceSchema_ resembling a _faceValue_ each loop and then appending to a final _cube_ value we return.  

A final digram to show a structured version of what is finally generated is equivalent to the diagram at the top.  

## Conclusion  
This project was not really intended to be a how to program ... type of deal but more as a how to percieve your programming approaches in a way. You see the rubik's cube might seem impossible to the average person, but to those who know the algorithms, the keys, can percieve the world where the impossible is the most easily attainable action. I will leave you with these last few words,   


**_dont be an insider_**

