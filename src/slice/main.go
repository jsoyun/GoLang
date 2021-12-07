package main

import "fmt"

func main(){
// a:= make([]int,0,8)
// 	fmt.Printf("len(a) = %d\n",len(a))
// 	fmt.Printf("cap(a)=%d\n",cap(a))

// 	a = append(a, 9)

// 	fmt.Printf("len(a) = %d\n",len(a))
// 	fmt.Printf("cap(a)=%d\n",cap(a))


a:= []int{1,2}
b:= append(a,3)
fmt.Printf("%p %p\n",a,b)
for i := 0; i<len(a); i++{
	fmt.Printf("%d,",a[i])

}
fmt.Println()
for i := 0; i<len(b); i++{
	fmt.Printf("%d,",b[i])
}
fmt.Println()

fmt.Println(cap(a)," ", cap(b))
}