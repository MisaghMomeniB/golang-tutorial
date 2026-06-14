package main

import "fmt"

// func main() {
// 	age := 20

// 	if age >= 18 {
// 		fmt.Println("Adult")
// 	}
// }

// func main() {
// 	age := 16

// 	if age >= 18 {
// 		fmt.Println("Adult")
// 	} else {
// 		fmt.Println("Under Age")
// 	}
// }

func main() {
	score := 85

	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Failed")
	}
}
