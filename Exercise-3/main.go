package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("not enough args")
	}
	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		t := strings.ReplaceAll(scan.Text(), old, new)
		fmt.Println(t)
	}
}
