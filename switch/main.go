package main

import (
	"fmt"
)

func main() {
	fmt.Println("Switch Statements in the Go Lang.")

	// finger := 4
	// fmt.Printf("Finger %d is ", finger)
	// switch finger {
	// case 1:
	// 	fmt.Println("Thumb")
	// case 2:
	// 	fmt.Println("Index")
	// case 3:
	// 	fmt.Println("Middle")
	// case 4:
	// 	fmt.Println("Ring")
	// case 5:
	// 	fmt.Println("Pinky")
	// default:
	// 	fmt.Println("Incorrect finger number.")
	// }

	// Multi Expression switch case
	// letter := "i"
	// switch letter {
	// case "a", "e", "i", "o", "u":
	// 	fmt.Printf("%s is a vowel", letter)
	// default:
	// 	fmt.Printf("%s is not vowel", letter)
	// }

	// Expression Less Switch (Evaluating to true than execute other wise next)
	// hour := 15 // hour in 24 hour format
	// // Using switch to determine the work shift
	// switch {
	// case hour >= 6 && hour < 12:
	// 	fmt.Println("It's the morning shift.")
	// case hour >= 12 && hour < 17:
	// 	fmt.Println("It's the afternoon shift.")
	// case hour >= 17 && hour < 21:
	// 	fmt.Println("It's the evening shift.")
	// case (hour >= 21 && hour <= 24) || (hour >= 0 && hour < 6):
	// 	fmt.Println("It's the night shift.")
	// default:
	// 	fmt.Println("Invalid hour.")
	// }

	/* Fallthrough
		Switch and case expressions need not be only constants. They can be evaluated at runtime too. In the program above num is initialized to the return value of the function number() in line no. 13. The control comes inside the switch and the cases are evaluated. case num < 100: in line no. 17 is true and the program prints 75 is lesser than 100. The next statement is fallthrough. When fallthrough is encountered the control moves to the first statement of the next case and also prints 75 is lesser than 200. 

		fallthrough should be the last statement in a case. If it is present somewhere in the middle, the compiler will complain that fallthrough statement out of place.


		Fallthrough happens even when the case evaluates to false

		So be sure that you understand what you are doing when using fallthrough. (As fallthrough tells the next case to be execute their block even the condition is false)

		One more thing is fallthrough cannot be used in the last case of a switch since there are no more cases to fallthrough. If fallthrough is present in the last case, it will result in the following compilation error cannot fallthrough final case in switch

	 */

	switch num := number(); { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
}

func number() int {
	num := 15 * 5
	return num
}
