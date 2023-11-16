package main

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	//input
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	//stage 1

	dataChanel := sliceToChannel(nums)

	//stage 3

	finalChannel := sq(dataChanel)

}
