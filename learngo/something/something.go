package something

import "fmt"

func SayHello() { //이건 public : 첫 글자 대문자
	fmt.Println("hello")
}

func sayBye() { //이건 private
	fmt.Println("bye")
}
