package main

import "fmt"


func main(){
	a,b:=fun1(1,3)
	fmt.Println(a,b)
}
func fun1(e,t int)(int,int){
	fun2(e,t)
	return e,t

}

func fun2(e,t int){
	fmt.Println("func2",e,t)
}
// func add(x int, y int) int{
// 	return x+y


// }
// func main(){ 
// 	for i:=0; i<10; i++{
// 		fmt.Printf("%d+%d=%d\n",i,i+2,add(i,i+2))

// 	}

	
// }