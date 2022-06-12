package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	pointersPoint()
	structsStructure()
	arraysArrange()
	slicesSlice()
	slicesLenCap()
	slicesMake()
	sliceTacToe()
	sliceAppend()
	rangeSimp()
	mapsMap()
	mapsMutate()
	functionValues()
	functionClosures()
}

func pointersPoint() {
	// go has pointers. points to a memory location of a variable. type of *T
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// & generates a pointer based on the operand
	// * denotes pointer's underlying value
}

func structsStructure() {
	// struct is just a collection of fields
	type Vertex struct {
		X int
		Y int
	}
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	v2 := Vertex{1, 2}
	p := &v2
	// we do not need the * for pointers and struct fields. convenience. not an error
	p.X = 1e9
	fmt.Println(v2)

	var (
		v1s = Vertex{1, 2}  // has type Vertex
		v2s = Vertex{X: 1}  // Y:0 is implicit
		v3s = Vertex{}      // X:0 and Y:0
		ps  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1s, ps, v2s, v3s)
}

func arraysArrange() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// arrays cannot be resized in go but slices can
}

func slicesSlice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// a slice of elements 1-3 from primes
	var s []int = primes[1:4]
	fmt.Println(s)

	// a slice is a flexible view into the elements of any array
	// slices do not store data. they are just references
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	/*
		A slice literal is like an array literal without the length.
		This is an array literal:
		[3]bool{true, true, false}
		And this creates the same array as above, then builds a slice that references it:
		[]bool{true, true, false}
	*/
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	sliteral := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(sliteral)

	sb := []int{2, 3, 5, 7, 11, 13}
	sb = sb[1:4]
	fmt.Println(sb)
	sb = sb[:2] // ie sb[0:2]
	fmt.Println(sb)
	sb = sb[1:] // ie sb[1:2]
	fmt.Println(sb)
}

func slicesLenCap() {
	// slice length = number of elements
	// slice capacity = number of elements of slice's array, counting from first element in slice
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	// zero value (default) of slice is nil
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func slicesMake() {
	/*
		Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
		The make function allocates a zeroed array and returns a slice that refers to that array:
		To specify a capacity, pass a third argument to make:
	*/
	a := make([]int, 5)
	printSliceOther("a", a)

	b := make([]int, 0, 5)
	printSliceOther("b", b)

	c := b[:2]
	printSliceOther("c", c)

	d := c[2:5]
	printSliceOther("d", d)
}

func printSliceOther(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func sliceTacToe() {
	// you can have a slice of a slice
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func sliceAppend() {
	// note s is a slice not an array. arrays have fixed size [0, [1], etc.
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func rangeSimp() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for idx, val := range pow {
		fmt.Printf("2**%d = %d\n", idx, val)
	}

	// range works for slices and maps

	// can omit idx or val as needed
	pow2 := make([]int, 10)
	for i := range pow2 { // only index
		pow2[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}

func mapsMap() {
	// maps map keys and values. same as in other languages. aka dictionary
	// maps can be nil. nil maps have no keys and cannot get keys added

	type Vertex struct {
		Lat, Long float64
	}

	var m map[string]Vertex

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// map literals are like struct literals but keys are required
	var mlit = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(mlit)
}

func mapsMutate() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] // test keys are present with two value assignment
	fmt.Println("The value:", v, "Present?", ok)
}

func functionValues() {
	// in go functions are values and can be passed around as such
	// functions can be arguments or return values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))    // ie the hypotenuse of 3 and 4
	fmt.Println(compute(math.Pow)) // ie 3 to the 4th power
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func functionClosures() {
	/*
		Go functions may be closures. A closure is a function value that references variables from outside its body.
		The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
		For example, the adder function returns a closure. Each closure is bound to its own sum variable
	*/
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
