package main

import "fmt"

//슬라이스로 만든 스택과 큐
func main() {
	stack := []int{}
	for i:= 1; i<= 5; i++ {
		stack=append(stack,i)
	}

    fmt.Println(stack)
	//스택의 길이가 0보다 클때까지 실행시켜
	for len (stack) >0 {
    var last int
	//=은 대입연산? 
	// 맨뒤에서 빼니까, 마지막,은 이거 /스택은 처읍무터 시작해서 맨마지막뺀거까지 
	last,stack = stack[len(stack)-1] , stack[:len(stack)-1]
	fmt.Println(last)

	}

	queue := []int{}
	for i:=1; i<=5; i++{
		queue = append(queue, i)

	}
	fmt.Println(queue)
	for len(queue) >0{
		var front int
		              //맨첫번째, 앞에서부터 빼니까
		front,queue = queue[0],queue[1:]
		fmt.Println(front)
	}

}