package main

import (
	//프런터
	"fmt"
	//한줄 입력했을때 찌꺼기 제거위해
	"strconv"
	//문자열을 숫자로 바꿔주는

	"strings"

	//표준입력위해
	"os"
	//한줄입력하기 위해
	"bufio"
)


func main() {
	fmt.Println("숫자를 입력하세요")
	reader := bufio.NewReader(os.Stdin)
	line,_ := reader.ReadString('\n')
	line= strings.TrimSpace(line)

	n1, _ := strconv.Atoi(line)

	line,_ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n2,_ := strconv.Atoi(line) 

	fmt.Printf("입력하신 숫자는 %d,%d입니다",n1,n2)
	fmt.Println("연산자를 입력하세요")

	line,_ =reader.ReadString('\n')
	line = strings.TrimSpace(line)

	// if line == "+" {
	// 	fmt.Printf("%d+%d=%d",n1,n2,n1+n2)
	// } else if line == "-" {
	// 	fmt.Printf("%d-%d=%d",n1,n2,n1-n2)
	// } else if line == "*"{
	// 	fmt.Printf("%d*%d=%d",n1,n2,n1*n2)
	// }else if line == "/" {
	// 	fmt.Printf("%d/%d=%d",n1,n2,n1/n2)
	// } else {
	// 	fmt.Println("잘못입력했음")
	// }

	//if문 switch로 
	switch line {
	case "+": 
	fmt.Printf("%d+%d=%d",n1,n2,n1+n2)
	case "-": 
	fmt.Printf("%d-%d=%d",n1,n2,n1-n2)
	case "*": 
	fmt.Printf("%d*%d=%d",n1,n2,n1*n2)
	case "/": 
	fmt.Printf("%d/%d=%d",n1,n2,n1/n2)
	default:
	fmt.Println("잘못입력하셨습니다")
	}
}