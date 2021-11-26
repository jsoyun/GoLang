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
	// for i := 0; i<5; i++ {
	// 	for j := 0; j< i+1; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	//별찍기 거꾸로 이건 내가한 방법1
	// for i:=0; i<5; i++{
	// 	for j:=5; 0<j-i ; j--{
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	//별찍기 거꾸로 답 방법2
	// for i:=0; i<5; i++{
	// 	for j:=0; j<5-i; j++{
	// 		fmt.Print("*")

	// 	}
	// 	fmt.Println()
	// }

	//별찍기 오른쪽 화살표같은삼각형

	// for i:=0; i<3; i++{
	// 	for k:=0; k<i+1; k++{
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()

	// }

	// for i := 0; i<2 ; i++ {
	// 	// fmt.Print("*")
	// 	for j:=0; i+j<2; j++{
	// 		fmt.Print("*")

	// 	}
	// 	fmt.Println()

	// }

	//별찍기 피라미드
	for i := 0; i < 4; i++ {

		for j := 0; j < 3-i; j++ {
			fmt.Print(" ")
		}
		for j:=0; j<2*i+1 ; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}

}