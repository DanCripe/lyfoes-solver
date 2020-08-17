package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		panic("must specify file")
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic("error reading file")
	}

	pot := 100

	var count int
	for {
		count++
		b := NewBoard(string(data))
		b.Solve()
		// b.Print()
		if b.Solved() {
			b.Print()
			fmt.Printf("after %d iterations\n", count)
			os.Exit(0)
		}
		if count == pot {
			fmt.Printf("%d %s\n", count, time.Now().String())
			pot *= 10
		}
	}
}
