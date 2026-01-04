package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}

	// for i := range len(arr) {
	// 	fmt.Printf("%v " ,arr[i]);
	// }

	for idx, val := range arr {
		fmt.Printf("Index:%v & Value:%v \n", idx, val);
	}

	fmt.Printf("\n%T \n", arr)

	// Dynamic Sized Arrays
	b := [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

	// If you specify the index with :, the elements in between will be zeroed.
    b = [...]int{100, 3: 400, 500} // b[3] = 400 and elements before it will be 0
    fmt.Println("idx:", b)  // idx: [100 0 0 400 500]



	// Arrays are VALUE types ⚠️ (Very Important):- When you assign or pass an array, a copy is made. 
	a := [3]int{1, 2, 3}
	c := a
	c[0] = 100

	fmt.Println(a) // [1 2 3]
	fmt.Println(c) // [100 2 3]

	// Use pointer if you want modification
	a_new := [3]int{1, 2, 3}
	update(&a_new)
	fmt.Println(a_new)  // [99 2 3]
}

func update(arr *[3]int) {
	arr[0] = 99
}
