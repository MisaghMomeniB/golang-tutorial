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

// func main() {
// 	score := 85

// 	if score >= 90 {
// 		fmt.Println("Grade A")
// 	} else if score >= 80 {
// 		fmt.Println("Grade B")
// 	} else if score >= 70 {
// 		fmt.Println("Grade C")
// 	} else {
// 		fmt.Println("Failed")
// 	}
// }

// func main() {
// 	age := 25
// 	hasLicense := true

// 	if age >= 18 && hasLicense {
// 		fmt.Println("Can Drive")
// 	}
// }

// func main() {
// 	isAdmin := false
// 	isOwner := true

// 	if isAdmin || isOwner {
// 		fmt.Println("Access Granted")
// 	}
// }

// func main() {
// 	isBlocked := false

// 	if !isBlocked {
// 		fmt.Println("User Active")
// 	}
// }

// func main() {

// 	if age := 20; age >= 18 {
// 		fmt.Println("Adult")
// 	}

// }

func main() {

	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of Week")

	case "Friday":
		fmt.Println("Weekend Coming")

	default:
		fmt.Println("Normal Day")
	}

}
