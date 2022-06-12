package main

import (
	"fmt"
	"image"
	"io"
	"math"
	"strconv"
	"strings"
)

type Vertex struct {
	X, Y float64
}

// you can only create methods for types defined in the same package
// you cannot create methdos for built in types. so not like kotlin extension functions then

func main() {
	// go does not have classes. you can define methods on a type though
	// seems like kotlin extension functions
	methodsExt()
	interfaceEx()
	emptyInterface()
	typeAssertion()
	typeSwitch()
	stringers()
	errorsErr()
	readersRead()
	imagineImages()
	genericTypeParams()
}

func methodsExt() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())
}

// method (kinda like an extension function)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
pointer receiver. ie you can create methods on *T receivers
Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
changing to v Vertex would have a different result b/c value receivers would operate on a copy of V. not V itself
*/
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*
functions with a pointer argument must take a pointer:

var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is
called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the
Scale method has a pointer receiver.
*/

/*
There are two reasons to use a pointer receiver.
The first is so that the method can modify the value that its receiver points to.
The second is to avoid copying the value on each method call. This can be more efficient
if the receiver is a large struct, for example.
*/

// -------
/*
An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods
*/
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func interfaceEx() {
	var i I = T{"hello"}
	i.M()
}

// interface{} is an empty interface. it can hold values of any type
// since every type implements at least 0 methods
func emptyInterface() {
	var i interface{}
	describe(i)
	// note interfaces will populate default values with nil if need be
	// will just print nil. no null pointer exception or the like

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// A type assertion provides access to an interface value's underlying concrete value.
func typeAssertion() {
	var i interface{} = "hello"

	// type assertion of the form    t := i.(T)
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) // would panic due to interface conversion
	fmt.Println(f)
}

func typeSwitch() {
	do(21)
	do("hello")
	do(true)
}

// you can have switch statements that use type to determine branch execution
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// One of the most ubiquitous interfaces is Stringer defined by the fmt package
// A Stringer is a type that can describe itself as a string. The fmt package (and many others)
// look for this interface to print values
func stringers() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// Go programs express error state with error values.
// The error type is a built-in interface similar to fmt.Stringer
func errorsErr() {
	i, err := strconv.Atoi("42gfd") // would work for "42"
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)
	// Functions often return an error value, and calling code should handle errors by
	// testing whether the error equals nil
}

// The io package specifies the io.Reader interface, which represents the read end of a stream of data
func readersRead() {
	/*
		Read populates the given byte slice with data and returns the number of bytes populated and an error value.
		It returns an io.EOF error when the stream ends.
		The example code creates a strings.Reader and consumes its output 8 bytes at a time
	*/
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// Package image defines the Image interface
func imagineImages() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

func genericTypeParams() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}

// Go functions can be written to work on multiple types using type parameters.
// The type parameters of a function appear between brackets, before the function's arguments
// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

//In addition to generic functions, Go also supports generic types. A type can be parameterized with a
// type parameter, which could be useful for implementing generic data structures
// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}
