package main

import "fmt"

//피보나치 재귀호출로~ 이렇게 수학정의 된거는 재귀호출이 쉬운데 반복수행은 반복문이더 쉬워
//재귀 호출: 함수 호출 과정(이 반복문보다 단계많음 느려), 인자 기록하고 (IP-> 시작포인트)
func main() {
	rst := f(10)
	fmt.Println(rst)
}

func f(x int) int {
	if x == 0{
		 return 1
	}
	if x ==1 {
		return 1
	}
	return f(x-1) + f(x-2)

}
