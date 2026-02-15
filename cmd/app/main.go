package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(
		"Number of CPUs: ",
		runtime.NumCPU(),
	)

	fmt.Println(
		"Number of Processors: ",
		runtime.GOMAXPROCS(5),
	)

	fmt.Println(
		"Number of Processors: ",
		runtime.GOMAXPROCS(5),
	)

	time.Sleep(1 * time.Minute)
}
