package main

import (
	"fmt"
	"math/rand"
	"time"
) 

type Result struct {
	strikes int
	balls int
}

func main() {
	rand.Seed(time.Now().UnixNano())
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
	//3개 배열 rst
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
				break
			}
		}
	}
	fmt.Println(rst)
	return rst
}
func InputNumbers() [3]int {
	//키보드로부터 0~9사이의 겹치지 않는 숫자3개를 입력받아 반환한다.
	var rst [3] int

	for{

        fmt.Println("겹치지 않는 0~9사이의 숫자 3개를 입력")
		//키보드로부터 입력받는
		var no int
		//사용하지 않는 변수는 밑줄
		_, err := fmt.Scanf("%d\n",&no)
		if err != nil {
			fmt.Println("잘못입력하셨습니다")
			//컨틴뉴로 포문 다시 시작
			continue
		}
		success := true
		     idx :=0
		     for no>0 { 
	                 n:= no %10
	                 no = no/10

					 duplicated := false

					 for j:=0; j<idx; j++{
						 //현재 숫자 i와 그다음 숫자j가 겹치면 
						 if rst[j]==n {
							 //겹치면 다시 뽑는다
							 duplicated = true
							 break
			 
						 }
			 
			 
					 }
					 if duplicated  {
						 fmt.Println("숫자가 겹치지 않아야한다")
						 success = false
						 break

					 }


					if idx >= 3{
						fmt.Println("3개보다 많은 숫자를 입력하셨습니다")
					    success= false
						break
					}


					 //인덱스 3개 배열안에
	                rst[idx] = n
	                  idx++
              }

			  if idx <3 {
				  fmt.Println("3개의 숫자를 입력하세요")
				  success = false
			  }

					if !success{
						continue

					}
					break

			}
			rst[0],rst[2] = rst[2],rst[0]
			fmt.Println(rst)
			return rst

}

func CompareNumbers(numbers,inputNumbers [3] int) Result {
	//두개의 숫자 3개를 비교해서 결과를 반환한다
	strikes:=0
	balls := 0
	for i := 0; i<3; i++ {
		for j:=0; j<3; j++ {
			if numbers[i]==inputNumbers[j]{
				if i== j {
					strikes++
				} else {
					balls++
				}
				break
			}
	}
	}
	return Result{strikes,balls}
}

func PrintResult(result Result){
	fmt.Printf("%dS%dB\n", result.strikes,result.balls)
}

func IsGameEnd(result Result) bool {
	//비교 결과가 3 스트라이크인지 확인
	return result.strikes ==3
}