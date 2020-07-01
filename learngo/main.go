package main

import (
	"fmt"
	"strings"
)

func test() {
	const n1 string = "kwon" //상수
	// var n2 string = "rnjs"   //변수
	// name := "jun" //변수 축약형 : 함수안에서만 가능

	fmt.Println(multiply(1, 2))
	// len, upper := lenAndUpper("kwon") //변수 두개 리턴
	l1, _ := lenAndUpper("jun") //하나 무시 가능
	fmt.Println(l1)
}

func multiply(a, b int) int {
	return a * b
}

//이렇게 리턴 타입 지정해도됨
func nakedReturn(name string) (length int, upper string) {
	length = len(name)
	upper = strings.ToUpper(name)
	return
}

//여러개 리턴가능
func lenAndUpper(name string) (int, string) {
	defer fmt.Println("I'm done :: return 되고나서 실행됨")
	return len(name), strings.ToUpper(name)
}

//여러 인자를 받을 때
func repeat(words ...string) {
	fmt.Println(words)
}

//반복문(for)
func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	/* 기본 형태 */
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return total
}

func canIDrink(age int) bool {
	//if문 안에 변수 선언가능
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func switchPrac(age int) bool {
	/* 일반 switch처럼 사용가능 */
	// switch koreanAge := age + 2; koreanAge {
	// case 10:
	// 	return false
	// case 18:
	// 	return true
	// }
	/* switch에 조건문도 작성가능*/
	switch {
	case age < 18:
		return false
	case age >= 18:
		return true
	default:
		return true
	}
}

//포인터 연산 가능, 주소 값
func mem() {
	a := 2
	b := &a
	*b = 20
	fmt.Println(a, *b)
}

// 배열
func arr() {
	names := []string{"a", "b", "c", "d", "e"}
	names = append(names, "f") //append는 원본은 수정하지 않음 : js의 slice와 같음
	fmt.Println(names)
}

func mapPrac() {
	//map :: key, value 정의
	rnjs := map[string]string{"name": "kwon", "age": "1"}
	for key, value := range rnjs {
		fmt.Println(key, value)
	}
}

//class 대신에 struct를 정의해서 활용
type person struct {
	name    string
	age     int
	favFood []string
}

func structPrac() {
	favFood := []string{"ramen", "chicken"}
	rnjs := person{name: "kwon", age: 26, favFood: favFood}
	fmt.Println(rnjs)
	fmt.Println(rnjs.name)
}

func main() {
	structPrac()
}
