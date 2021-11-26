package main

import "fmt"

func main() {


	// //3X2빼고 구구단
	// for dan := 1; dan <= 9; dan++ {
		

	// 	fmt.Printf("%d단\n", dan)

	// 	for j := 1; j <= 9; j++ {

	// 		if dan==3&&j ==2 {
	// 			continue
	// 		}
	// 		fmt.Printf("%d*%d=%d\n", dan, j, dan*j)
	// 	}
	// 	fmt.Println()
	// }

	//별찍기
	for i := 0; i<5; i++ {
		for j := 0; j< i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}


}