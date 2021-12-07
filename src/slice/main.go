package main

import "fmt"

func main() {
	// a:= make([]int,0,8)
	// 	fmt.Printf("len(a) = %d\n",len(a))
	// 	fmt.Printf("cap(a)=%d\n",cap(a))

	// 	a = append(a, 9)

	// 	fmt.Printf("len(a) = %d\n",len(a))
	// 	fmt.Printf("cap(a)=%d\n",cap(a))

	//////////////////////////////////len cap 메모리를 새로 확보해야돼서 주소도 달라짐  새로운 메모리에 담아서
	// a:= []int{1,2}
	// b:= append(a,3)
	// fmt.Printf("%p %p\n",a,b)
	// for i := 0; i<len(a); i++{
	// 	fmt.Printf("%d,",a[i])

	// }
	// fmt.Println()
	// for i := 0; i<len(b); i++{
	// 	fmt.Printf("%d,",b[i])
	// }
	// fmt.Println()

	// fmt.Println(cap(a)," ", cap(b))

	///////make로 넣기//메모리 같아서 b값 바꾸니까 a도 바뀜

	// a:= make([]int, 2,4)
	// a[0]=1
	// a[1]=2
	// b:= append(a,3)

	// fmt.Printf("%p %p\n",a,b)
	// fmt.Println(a)
	// fmt.Println(b)
	// b[0]=4
	// b[1]=5

	// fmt.Println(a)
	// fmt.Println(b)
	//////처음부터 아예 공간을 다르게 확보하게 싶다면
	// a:= []int{1,2}
	a := make([]int, 2, 4)
	a[0] = 1
	a[1] = 2
	//슬라이스 새로 만듦 
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}
	b = append(b, 3)
	fmt.Printf("%p %p\n",a,b)
}