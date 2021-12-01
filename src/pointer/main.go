// package main

// import "fmt"

// func main(){
//    //a는 int 형 이라 선언
// 	var a int
// 	//p는 int형 포인트
// 	var p *int
// 	//p에 a 주소 대입, p의 값은 a의 주소를 가르키고 있다
// 	p = &a
// 	//a는 3으로 대입
// 	a=3
// 	//3
// 	fmt.Println(a)
// 	//0xc000100028 메모리의 주소 나옴
// 	fmt.Println(p)
// 	//포인터가 가르키는 값 3(주소가 가르키는 값)
// 	fmt.Println(*p)

////
// //a는 int 형 이라 선언
// var a int
// var b int
// //p는 int형 포인트
// var p *int
// //p에 a 주소 대입, p의 값은 a의 주소를 가르키고 있다
// p = &a
// //a는 3으로 대입
// a=3
// b=2
// //3
// fmt.Println(a)
// //0xc000100028 메모리의 주소 나옴
// fmt.Println(p)
// //포인터가 가르키는 값 3(주소가 가르키는 값)
// fmt.Println(*p)

// p= &b
// fmt.Println(b)
// //0xc000100028 메모리의 주소 나옴
// fmt.Println(p)
// //포인터가 가르키는 값 3(주소가 가르키는 값)
// fmt.Println(*p)

// //포인터가 왜 사용되는지 보자`~~~~`
// var a int
//   a=1
//   //포인터형 *x으로 해서 a의 주소값으로 해야함
//   Increase(&a)

//   fmt.Println(a)
// }

// func Increase(x *int) {
// 	//x가 가지고 있는 값은 메모리 주소인데
// 	//그 메모리가 가르키는 주소의 값은 = 메모리가 가르키는 주소의 (원래)값 + 1
// *x=*x+1

// //얘는 아까 x++와 같은 거라 볼 수 있쥐, 그때 포인터로 안하면 a에서 값만 복사돼서 x++여도 x는 함수 안이라서 a가 늘어나진 않았어
// }

package main

import "fmt"


type Student struct {
name string
age int
grade string
class string

}

func (s *Student) PrintSungjuk(){

	fmt.Println(s.class,s.grade)
}

func (s *Student) InputSungjuk(class string, grade string){
    s.class = class
	s.grade =grade
}


func main(){
	var s Student= Student{name: "t",age:23, class: "수학",grade: "A"}
 s.InputSungjuk("과학","c")
 s.PrintSungjuk();
}