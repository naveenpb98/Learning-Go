package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var sum int
	var val int
	for {
		_, err := fmt.Fscanln(os.Stdin, &val)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		sum += val
	}
	fmt.Println(sum)
}

// To generate EOF use Ctrl + D
