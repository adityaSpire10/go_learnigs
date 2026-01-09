package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in Go Lang")

	// for i := 1; i <= 10; i++ {

	// 	// if i > 5 {
	// 	// 	break; //Loop is terminated if i > 5
	// 	// }

	// 	if i%2==0 {
	// 		continue; //Loop is terminated if i > 5
	// 	}

	// 	fmt.Printf("%d ", i);
	// }




	// Nested for loop
	// n := 5
	// for i := range n {
	// 	for j := 0; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }






	// 	Labels
	// Labels can be used to break the outer loop from inside the inner for loop. Let’s understand what I mean by using a simple example.

	// We want this loop to be stop when i and j both are equal to each other
	for i := range 3 {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				break
			}
		}
	}
	// In the program above, I have added a break inside the inner for loop when i and j are equal. This will break only from the inner for loop and the outer loop will continue. This program will print.
	// i = 0 , j = 1
	// i = 0 , j = 2
	// i = 0 , j = 3
	// i = 1 , j = 1
	// i = 2 , j = 1
	// i = 2 , j = 2
	//This is not the intended output. We need to stop printing when both i and j are equal i.e when they are equal to 1.

	// This is where labels come to our rescue. A label can be used to break from an outer loop. Let’s rewrite the program above using labels.
	fmt.Println("Labels Example:")
outer:
	for i := range 3 {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				break outer
			}
		}
	}


	
	// Multiple initializations
	fmt.Println("Multiple initialization in the loop")
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

}
