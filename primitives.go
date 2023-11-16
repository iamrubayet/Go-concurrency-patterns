// this is a synchornus code and program will be blocked for that somefunc to be executed
// package main

// import "fmt"

// func someFunc(num string) {
// 	fmt.Println(num)
// }

// func main() {
// 	someFunc("2")

// 	fmt.Println("hi")

// }

// asynchronus code or goroutine

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func someFunc(num string) {
// 	fmt.Println(num)
// }

// func main() {
// 	go someFunc("1")
// 	go someFunc("2")
// 	go someFunc("3")

// 	time.Sleep(time.Second * 2) // blocking point or join point

// 	fmt.Println("hi")

// }

// ///////////////////////////////////go channels

// package main

// import "fmt"

// func main() {
// 	myChannel := make(chan string)

// 	go func() {
// 		myChannel <- "data"
// 	}()

// 	msg := <-myChannel // blocking point or join point

// 	fmt.Println(msg)
// }

// select statement

// package main

// import "fmt"

// func main() {
// 	myChannel := make(chan string)
// 	anotherChannel := make(chan string)

// 	go func() {
// 		myChannel <- "data"
// 	}()

// 	go func() {
// 		anotherChannel <- "data222222"
// 	}()

// 	select {
// 	case msgfrommychannel := <-myChannel:
// 		fmt.Println(msgfrommychannel)
// 	case msgfromanotherchannel := <-anotherChannel:
// 		fmt.Println(msgfromanotherchannel)
// 	}

// }
