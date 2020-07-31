package main

import (
	"fmt"
)

func loopData(data []int, handleData chan<- int) {
	defer close(handleData)
	for i := range data {
		handleData <- data[i]
	}
}
func main() {
	data := []int{1, 2, 3, 4}

	handleData := make(chan int)
	go loopData(data, handleData)

	for num := range handleData {
		fmt.Println(num)
	}

}
