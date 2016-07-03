/*

Copyright David Gee 2016

License
View license here https://creativecommons.org/licenses/by/4.0/

This code exercises creating types, methods on those types and wrapping functions, also pointers.

*/

package main

import (
	"fmt"
	"io"
	"strings"
)

// Types

type WriteFunc func(p []byte) (n int, err error)

type interfaceWrite interface {
	Write(p []byte) (n int, err error)
}

// Funcs

func (wf WriteFunc) Write(p []byte) (n int, err error) {
	return wf(p)
}

func myWrite(p []byte) (n int, err error) {
	fmt.Printf("Byte array = %v, String = %v\n", p, string(p))
	// returns length of p (n int) and the err of type error, which is nil
	return len(p), nil
}

// Main entry()

func main() {
	// This example uses io.Copy and the Write() invocation
	fmt.Println("Exercise [1] ==== Variable invocation of WriteFunc(myWrite)")
	fmt.Println(io.Copy(WriteFunc(myWrite), strings.NewReader("Hello")))
	fmt.Println("Exercise [1] ==== END \n\n")

	// Create a variable based on the type WriteFunc
	// and fill it with custom function myWrite()
	var a WriteFunc
	a = myWrite

	// next, we call the Write method on the type WriteFunc, which happens to point at myWrite
	fmt.Println("Exercise [2] ==== Write method invocation on the var a which is of type WriteFunc, containing myWrite")
	fmt.Println(a.Write([]byte("Hello")))
	fmt.Println("Exercise [2] ==== END \n\n")

	// Here we invoke Write directly on the type, wrapping myWrite as we do it
	fmt.Println("Exercise [3] ==== Write method invocation on the WriteFunc type, being passed the func myWrite directly")
	fmt.Println(WriteFunc(myWrite).Write([]byte("Hello")))
	fmt.Println("Exercise [3] ==== END \n\n")

	// Let's create an interface var. Note, you can do this as a pointer using new()
	var b interfaceWrite
	b = WriteFunc(myWrite)

	// We point it at WriteFunc with the myWrite custom func wrapped
	fmt.Println("Exercise [4] ==== Write method invocation on the interface typed variable of b")
	fmt.Println(b.Write([]byte("Hello")))
	fmt.Println("Exercise [4] ==== END \n\n")

	// Now let's do this with a pointer

	// Create pointer using new to WriteFunc type
	c := new(WriteFunc)
	// Point contents of c type WriteFunc to myWrite func
	*c = myWrite
	// Invoke Write method on c
	fmt.Println("Exercise [5] ==== Pointer based Write method invocation on the WriteFunc type, pointed at the myWrite func")
	fmt.Println(c.Write([]byte("Hello")))
	fmt.Println("Exercise [5] ==== END \n\n")

	/*
		DO THIS BIT LAST - Interesting outcome
		Let's do a pointer to an interface, and point that at the WriteFunc type passing the function myWrite
		When you've got the above working, uncomment below and run it.
	*/

	// d := new(interfaceWrite)
	// *d = WriteFunc(myWrite)
	// fmt.Println("Exercise [6] ==== Pointer based Write method invocation on the interfaceWriter interface, pointed at the wrapped myWrite func")
	// fmt.Println(d.Write([]byte("Hello")))
	//	fmt.Println("Exercise [6] ==== END \n")

	/*
		I knew this wouldn't work because I did some homework!

		Interface values are represented as a two-word pair giving a pointer to information
		about the type stored in the interface and a pointer to the associated data.

		Not sure if it's possible to separate these pointers? Not sure why you would want to either.

	*/

	// Next, let's set up a pointer as exercise 5 and create an interface to that pointer

	// Create pointer using new to WriteFunc type
	e := new(WriteFunc)
	// Point contents of c type WriteFunc to myWrite func
	*e = myWrite

	var ei interfaceWrite

	ei = e
	// Invoke Write method on c
	fmt.Println("Exercise [7] ==== Interface based method invocation on a pointer to the wrapped myWrite func via WriteFunc")
	fmt.Println(ei.Write([]byte("Hello")))
	fmt.Println("Exercise [7] ==== END \n\n")

}
