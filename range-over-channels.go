// range-over-channels https://books.studygolang.com/gobyexample/range-over-channels/
package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
