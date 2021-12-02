package main

import (
	"fmt"
	"math/rand"
) 


func main() {
	//무작위 숫자 3개를 만든다
	numbers := MakeNumbers()

	cnt := 0


	for{
		cnt++
	//사용자의 입력을 받는다. 
	//변수에 함수넣는다,대입한다
	inputNumbers := InputNumbers()
	//결과 비교
	result := CompareNumbers(numbers,inputNumbers)

	PrintResult(result)

	//3s인가?
	if IsGameEnd(result){

		//게임끝
		break
	}
	}
	//게임끝 몇번만에 맞췄는지 출력
	fmt.Printf("%d 번만에 맞췄습니다.\n", cnt)

}

func MakeNumbers()[3]int {
	//0~9사이의 겹치지 않는 무작위 숫자 3개를 반환한다.
	var rst [3] int

	for i :=0; i<3; i++ {
		//겹쳤을 경우 다시 뽑을 때돌아가려고 for문
		for{


			n:=rand.Intn(10)
			//겹치는지 안겹치는지에 대한 플래그 변수 만들기
			duplicated := false

			for j:=0; j<i; j++{
				//현재 숫자 i와 그다음 숫자j가 겹치면 
				if rst[j]==n {
					//겹치면 다시 뽑는다
					duplicated = true
					break
	
				}
	
	
			}
			if !duplicated {
				rst[i]=n
			}
		}
	}
	
	return rst
}
func InputNumbers() [3]int {
	//키보드로부터 0~9사이의 겹치지 않는 숫자3개를 입력받아 반환한다.
	var rst [3] int
	return rst

}

func CompareNumbers(numbers,inputNumbers [3] int) bool {
	//두개의 숫자 3개를 비교해서 결과를 반환한다
	return true
}

func PrintResult(result bool){
	fmt.Println(result)
}

func IsGameEnd(result bool) bool {
	//비교 결과가 3 스트라이크인지 확인
	return result
}