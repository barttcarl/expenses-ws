package main

import (
		"fmt"
		"time"
		"math"
		"math/rand"
		"math/cmplx"
		"github.com/alexbeca/expenses-ws/stringutil"
)

/* 
Comparing with methods in Java we will see there are not keywords public, private, static at the begin of the functions in Go

the following function will be in Java like this:
	public Integer add(Integer x, Integer y){ return Integer }
*/
func add(x int, y int) int {
	return x + y
}

/*
When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
*/
func addV2(x, y int) int {
	return x + y
}

/*
A function can return any number of results.
*/
func swap(x, y string) (string , string) {
	return y, x
}

/*
Go's return values may be named. If so, they are treated as variables defined at the top of the function.
These names should be used to document the meaning of the return values.
A return statement without arguments returns the named return values. This is known as a "naked" return.
Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
*/
func split(sum int) (x, y int){
	x = sum * 4 / 9
	y = sum - x
	return
}

/*
The var statement declares a list of variables; as in function argument lists, the type is last.

A var statement can be at package or function level. We see both in this example.
*/

var python, java bool //Initial value is false

/*
Numeric Constants
Numeric constants are high-precision values.
*/
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1 (in Binary ..0010), or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}


func main() {
	fmt.Println("hello, world")
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is:", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(add(3, 4))
	fmt.Println(addV2(3, 19))

	/* 
	In Go, := is for declaration + assignment, whereas = is for assignment only.
	For example, var foo int = 10 is the same as foo := 10. 
	*/
	a, b := swap("hello", "devops")
	fmt.Println(a, b)

	fmt.Println(split(10))

	var idx int // Initial value is 0
	fmt.Println(python, java, idx)

	/*
	A var declaration can include initializers, one per variable.
	If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	*/
	var i, j int = 1, 2
	var golang, ruby, php = true, false, "no!"
	fmt.Println(i, j, golang, ruby, php)

	/*
	Basic types

	bool
	string
	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte // alias for uint8
	rune // alias for int32
		// represents a Unicode code point
	float32 float64
	complex64 complex128	
	*/

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	/*
	Variables declared without an explicit initial value are given their zero value.
	*/
	var f float64
	var s string
	fmt.Printf("%v %q\n", f, s)

	/*
	The expression T(v) converts the value v to the type T
	*/

	l := 42
	m := float64(l)
	n := uint(m)

	fmt.Println(l, m, n)

	/*
	Constants are declared like variables, but with the const keyword.

	Constants can be character, string, boolean, or numeric values.

	Constants cannot be declared using the := syntax.
	*/

	const Pi = 3.14
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	/* Constants exercise */
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}