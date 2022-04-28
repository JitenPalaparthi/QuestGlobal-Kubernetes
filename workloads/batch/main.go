package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	count := 0
	flag.IntVar(&count, "count", 100, "--count=100")
	flag.Parse()
	for i := 1; i <= count; i++ {
		fmt.Println("Current invoice is ->", i)
		time.Sleep(time.Second * 1)
	}
}
