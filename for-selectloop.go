// package main

// func main() {
// 	charChannell := make(chan string, 3)
// 	chars := []string{"a", "b", "c", "d", "e", "f", "g"}

// 	for _, s := range chars {
// 		select {
// 		case charChannell <- s:
// 		default:
// 			println("channel is full")
// 		}
// 	}
// 	close(charChannell)

// 	for result := range charChannell {
// 		println(result)
// 	}
// }

// infinte amount time

// package main

// import "time"

// func main() {

// 	go func() {
// 		for {
// 			select {
// 			default:
// 				println("doing work")
// 			}
// 		}

// 	}()

// 	time.Sleep(time.Second * 2)

// }
