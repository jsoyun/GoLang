// package main

// import "fmt"

// func main() {
// 	// //정수형 배열 열개짜리
// 	// var A [10] int
// 	// // for i :=0; i<10; i++ {
// 	//  //배열(A)의 길이(len) 이렇게 써도 됨
// 	// for i :=0; i<len(A); i++ {
// 	// 	A[i] =i*i

// 	// }
// 	// fmt.Println(A)

// 	// s := "hello 월드"
// 	// //한국어는 한글자당 3byte씩 그래서 배열 길이 12임
// 	// fmt.Println("len(s)=", len(s))
// 	// for i := 0; i < len(s); i++ {
// 	// 	fmt.Print(s[i],",")
// 	// }

// 	//string []byte로도 가능하지만 []rune으로도 나타낼 수 있음
// 	s := "hello 월드"
// 	s2 := []rune(s)
// 	//룬배열로 해서 길이가 8이됨
// 	fmt.Println("len(s2)=", len(s2))
// 	for i := 0; i < len(s2); i++ {
// 		fmt.Print(string(s2[i]),", ")
// 	}

// }

// //18강 배열복사
// package main

// import "fmt"

// func main(){
// 	arr := [5]int{1,2,3,4,5}
// 	clone := [5]int {}

// 	//0부터 길이 5까지 클론의 인덱스에 해당하는 arr 복사
// 	for i:=0; i<5; i++ {
// 		clone[i]= arr[i]

// 	}

// 	fmt.Println(clone)

// }

//18강 배열 역순으로 순서 바꾸기
//방법1. 임시배열에 순서 바꿔복사하고 원래 배열에 그걸 넣기
package main

import "fmt"

// func main(){
// 	arr := [5]int{1,2,3,4,5}
// 	temp := [5]int {}

// 	//0부터 길이 5까지 클론의 인덱스에 해당하는 arr 복사
// 	// for i:=0; i<5; i++ {
// 		for i:=0; i<len(arr); i++ {
// 		//4는 여기서 배열의 길이(len(arr))-1 임
// 		// temp[i]= arr[4-i]
// 		temp[i]= arr[len(arr)-1-i]

// 	}
// 	for i := 0; i<len(arr); i++ {
// 		arr[i]= temp[i]
// 	}

// 	fmt.Println(temp)
// 	fmt.Println("arr:", arr)

// }

//방법2. 자리만 바꿔서 역순으로 만들기
 func main(){
	 arr := [5] int {1,2,3,4,5}
	 for i:=0; i<len(arr)/2; i++ {

	//이중대입
	arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]

	 }
	 fmt.Println(arr)


 }