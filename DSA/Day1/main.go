package main

import (
	"fmt"
)

/*

//  ======================== 1. Constant Equation's Time Complexity Example : ==========================
func main() {
	x := 10							// 1 unit
	fmt.Println(x)					// 1 unit

	if x == 5 {
		fmt.Println("Hello")		// 1 unit
	} else {
		fmt.Println("Welcome")		// 1 unit
	}
}
//  So this program will run in 3 unit of time which is a constant so, time complexity of this program will be O(1).

*/





/*
//  ======================== 1. Linear Equation's Time Complexity Example : ==========================

func main () {
	x := 1					// 1 unit

	for ; x < 5; x++ {		// 4 unit (It depends on the x)
		fmt.Println(x)		// 4 unit (It depends on the x)
	}

	fmt.Println("Final value of x:", x) // 1 unit
}

// As the execution time is depending upon the variable x. So this is a linear equation so the time complexity will be O(n).
*/



/*
//  ======================== 1. Quadratic Equation's Time Complexity Example : ==========================

func main () {
	x := 2
	
	y := 3
	
	var runtime int = 0
	
	for i := 0; i < x; i++ {			// for range x
		for j := 0; j < y; j++ {		// for range y
			runtime = runtime + 1		
			fmt.Println(runtime)
		}
	}

	fmt.Printf("So the code executed for %v * %v : %v", x, y, runtime)
}
*/



/*
//  ======================== 1. Cubic Equation's Time Complexity Example : ==========================

func main () {
	x := 2
	y := 3
	z := 5
	
	var runtime int = 0
	
	for range x {
		for  range y {
			for range z {
				runtime = runtime + 1		
				fmt.Println(runtime)
			}
		}
	}

	fmt.Printf("So the code executed for %v * %v * %v : %v", x, y, z, runtime)
}
*/






//  ======================== 1. Logarithmic Equation's Time Complexity Example : ==========================

func main () {
	x := 10
	
	for x > 1 {
		fmt.Println("Hello")
		x = x/2
		fmt.Println("The value of x is: ",x)
	}
}
