package main

/*
Every Go program is made up of packages.
Programs start running in package main
*/

// vs code go extension tends to delete imports if you don't use them on save
// or it might be something within the language itself. apparently go projects
// won't build correctly with unused imports or variables

//import "fmt" for a 1 liner. these are package imports
import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

// these bools default to false. scoping works as expected
var c, python, java bool

// cantdo := 5. only works inside of functions

// const cannot use :=
const Pi = 3.14

// run the code as $ go run tour.go
// or do $ go build tour.go && ./tour
func main() {
	fmt.Println("hello world")
	fmt.Println("chinese or japanese characters. hello 世界")
	// those characters are actually chinese for "world". who would've known
	fmt.Println("the time is ", time.Now())
	fmt.Println("using a random number ", rand.Intn(10))

	// all exported names begin with a capital letter. ie. can't do println
	fmt.Println("calling a function", add(5, 7))

	swapa, swapb := swap("hello", "world")
	fmt.Println("swap function ", swapa, swapb)

	var idem int // defaults to 0
	fmt.Println(idem, c, python, java)

	// types can be omitted with initializers
	var c2, python2, java2 = true, false, "no!"
	// you can of course decalre variables 1 at a time too
	fmt.Println(c2, python2, java2)

	// := will infer the type of a var. must use inside a function
	kinfer := 3
	fmt.Println(kinfer)

	basicTypes()
	typeConversions()
	numericConstants()
	loopingThereCanOnlyBeOne()
	ifStmt()
	switchStmt()
	deferStmt()
}

// could also have params as     x, y int
func add(x int, y int) int {
	return x + y
}

// functions can return multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// a naked return. will return x and y as named results. discouraged for larger functions
// I probably won't ever use this then
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func basicTypes() {
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// there's also strings, runes, bytes, and normal ints
}

func typeConversions() {
	// The expression T(v) converts the value v to the type T
	i := 42
	f := float64(i)
	u := uint(f)
	// note using Printf not Println
	fmt.Printf("type conversions Type: %T Value: %v\n", f, f)
	fmt.Printf("type conversions Type: %T Value: %v\n", u, u)

	// go has type inference
	ii := 42           // int
	fi := 3.142        // float64
	gi := 0.867 + 0.5i // complex128
	fmt.Printf("type infer %T %v\n", ii, ii)
	fmt.Printf("type infer again %T %v\n", fi, fi)
	fmt.Printf("type infer again v2 %T %v\n", gi, gi)
}

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func numericConstants() {
	fmt.Print("\n \nNumeric constants \n \n")
	const (
		// Create a huge number by shifting a 1 bit left 100 places.
		// In other words, the binary number that is 1 followed by 100 zeroes.
		Big = 1 << 100
		// Shift it right again 99 places, so we end up with 1<<1, or 2.
		Small = Big >> 99
	)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func loopingThereCanOnlyBeOne() {
	// go only has for loops
	fmt.Print("\n \nFor loops \n \n")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// do not use () for the loops. actually an error

	// here is a for loop that behaves like a while
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// init and post parts are optional. note tool chain removes the ;
	// could be      for ; sum3 < 1000 ; {
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3)

	/* infinite loop is
		for {
	    }
	*/
}

func ifStmt() {
	// don't use (). similar to for loop
	if 1 < 2 {
		fmt.Println("1 is less than 2")
	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	if 2 < 1 {
	} else {
		fmt.Println("2 is not less than 1")
	}
}

// Like for, the if statement can start with a short statement to execute before the condition.
// Variables declared by the statement are only in scope until the end of the if
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func switchStmt() {
	/*
		Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case,
		not all the cases that follow.  In effect, the break statement that is needed at the end of each case in those
		languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants,
		and the values involved need not be integers
	*/
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// by the way go will try to autoimport the packages you need after saving

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// switch statements can omit the condition. ie just switch {}. is a switch true
}

func deferStmt() {
	// defer will execute after surronding function returns
	defer fmt.Println("world")
	fmt.Println("hello")

	// defered function calls are put on a stack. LIFO
	fmt.Println("counting")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
