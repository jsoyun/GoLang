package main

import "fmt"

//뒤에 하나씩 지워나가는 함수//배열 반환,int 맨뒤에 값도 반환
func RemoveBack(a []int) ([]int, int){
	//처음부터 a배열길이에서 1뺀것, 즉 맨뒤에를 하나 없앤다
	return a[ :len(a)-1], a[len(a)-1]

}
func RemoveFront(b []int) ([]int,int) {
	return b[1:],b[0]
}
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
	//
	// a := make([]int, 2, 4)
	// a[0] = 1
	// a[1] = 2
	// //슬라이스 새로 만듦
	// b := make([]int, len(a))
	// for i := 0; i < len(a); i++ {
	// 	b[i] = a[i]
	// }
	// b = append(b, 3)
	// fmt.Printf("%p %p\n",a,b)

	/////25강
	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// b := a[4:8]
	// // b:= []int{5,6,7,8}
	// // c:=a[4:]
	// // fmt.Println(b)
	// // fmt.Println(c)
	// b[0] = 1
	// b[1] = 2
	// fmt.Println(a)

	//하나씩 돌면서 뒤에 지움
	a:=[]int{1,2,3,4,5,6,7,8,9,10}

	for i :=0; i<5; i++{
		var lastlostone int
		a, lastlostone= RemoveBack(a)
       fmt.Printf("%d,",lastlostone)
	}
	fmt.Println()
	fmt.Println(a)

	b:=[]int{1,2,3,4,5,6,7,8,9,10}

	for i :=0; i<5; i++{
		var frontone int
		b, frontone =  RemoveFront(b)
       fmt.Printf("%d,",frontone )
	}

	fmt.Println()
	fmt.Println(b)
	

}