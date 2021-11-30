package main

import "fmt"

func main() {
	// //정수형 배열 열개짜리
	// var A [10] int
	// // for i :=0; i<10; i++ {
	//  //배열(A)의 길이(len) 이렇게 써도 됨
	// for i :=0; i<len(A); i++ {
	// 	A[i] =i*i

	// }
	// fmt.Println(A)


	// s := "hello 월드"
	// //한국어는 한글자당 3byte씩 그래서 배열 길이 12임
	// fmt.Println("len(s)=", len(s))
	// for i := 0; i < len(s); i++ {
	// 	fmt.Print(s[i],",")
	// }


	//string []byte로도 가능하지만 []rune으로도 나타낼 수 있음
	s := "hello 월드"
	s2 := []rune(s)
	//룬배열로 해서 길이가 8이됨
	fmt.Println("len(s2)=", len(s2))
	for i := 0; i < len(s2); i++ {
		fmt.Print(string(s2[i]),", ")
	}


}