package main

import (
	"fmt"
	"time"
)

//goroutine 를 쓰면 병렬처리 가능. 단 main이 살아있는 동안에만 실행
//goroutine 사용 후 그 값의 리턴은 채널을 이용해서 받아야 됨
func main() {
	c := make(chan string)
	people := [2]string{"rnjs", "kwon"}
	for _, person := range people {
		go count(person, c)
	}
	//main이 끝났지만 channel로부터 값을 받기위해 기다림 : `<-c` 사용
	//blocking operation
	for i := 0; i < len(people); i++ {
		fmt.Println("waiting for ", i)
		fmt.Println(<-c)
	}
}

func count(person string, c chan string) {
	time.Sleep(time.Second * 2)
	c <- person + " HI"
}
