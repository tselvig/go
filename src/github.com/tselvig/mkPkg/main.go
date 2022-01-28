package main

// Tanner Selvig

import (
	"fmt"

	"github.com/tselvig/myPkg/test1" // import package you created
)

func main() {
	out := test1.DoubleValue(8)   // Call function in your package
	fmt.Printf("out = %d\n", out) // should print "out = 16"
	// add call to TripleValue at this point
	outT := test1.TripleValue(20)
	fmt.Printf("out = %d\n", out)
}