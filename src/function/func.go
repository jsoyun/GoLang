// package main

// import "fmt"

// func main(){
// 	a,b:=fun1(1,3)
// 	fmt.Println(a,b)
// }
// func fun1(e,t int)(int,int){
// 	fun2(e,t)
// 	return e,t

// }

// func fun2(e,t int){
// 	fmt.Println("func2",e,t)
// }

// func add(x int, y int) int{
// 	return x+y

// }
// func main(){
// 	for i:=0; i<10; i++{
// 		fmt.Printf("%d+%d=%d\n",i,i+2,add(i,i+2))

// 	}

// }

//재귀호출만들시 무한되지 않게 하기 ex) return 쓰기

// package main

// import "fmt"

// func main(){
// 	f1(10)
// }

// func f1(x int){
// 	if x == 0{
// 		return
// 	}
// 	fmt.Printf("f1(%d) before call f1(%d)\n", x, x-1)
// 	f1(x-1)
// 	fmt.Printf("f1(%d) after call f1(%d)\n", x,x-1)
// }

//재귀호출사용해서 1에서 10까지 합계를 출력하는 프로그램

package main

import "fmt"
func main(){
	s:=sum(10,0)
	fmt.Println(s)
}

func sum(x int, s int) int {
	if x==0 {
		return s
	}
	s += x
	return 	sum(x-1,s)
}