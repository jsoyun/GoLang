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
// package main

// import "fmt"

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
//  func main(){
// 	 arr := [5] int {1,2,3,4,5}
// 	 for i:=0; i<len(arr)/2; i++ {

// 	//이중대입
// 	arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]

// 	 }
// 	 fmt.Println(arr)

//  }

// func main(){
// 	arr := [11] int {0,5,4,9,1,2,3,8,3,6,4 }
// 	temp := [10]int{}

// 	for i := 0; i < len(arr); i++ {
// 		idx := arr[i]
// 		temp[idx]++
// 	}
// //현재 인덱스 나타내는 idx
// idx := 0

// 	for i:=0; i<len(temp); i++{
// 		for j := 0; j <temp[i]; j++ {
// 			arr[idx] = i
// 			idx++

// 		}
// 	}
// 	fmt.Println(arr)

// }

//19강

// package main

// import "fmt"

// type Person struct {
//   name string
//   age int

// }
// //어떤 타입의 기능이냐 //p, 펄슨이 가지고 있는 기능 그 이름은 프린트 네임
// //괄호안에 비었어 입력값은 없다
// //펄슨에 메서드 ,기능 추가함
// func (p Person) PrintName() {
// 	fmt.Print(p.name)
// }

// func main() {
// var p Person
// p1 := Person{"개똥",15}
// p2 := Person{name:"rk똥",age:20}
// p3 := Person{name: "carson"}
// p4 := Person{}

// fmt.Println(p,p1,p2,p3,p4)

// p.name = "sss"
// p.age= 24

// fmt.Println(p)

// p.PrintName()
// }

package main

import "fmt"

type Student struct{
	name string
	class int

	sungjuk Sungjuk
}

type Sungjuk struct {
	name string
	grade string
}

func (s Student) ViewSungJuk(){
	fmt.Println(s.sungjuk)
}

func main() {

	var s Student
	s.name="철"
	s.class = 1

	s.sungjuk.name ="수학"
	s.sungjuk.grade= "A"

	s.ViewSungJuk()

}