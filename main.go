package main

import (
	"os/signal"
	"sync"
	"io/ioutil"
	"math/rand"
	"os"
	"syscall"
	"time"
)

var m sync.Mutex
var index int

func main() {
	if len(os.Args) < 2 {
		panic("must specify file")
	}
	data, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		panic("error reading file")
	}

	for i := 0; i < 4; i++ {
		go run(i, string(data))
	}
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGUSR1)
	for {
		<-sigch
		SetVerbose(true)
		time.Sleep(time.Second)
		SetVerbose(false)
	}
}

func run(i int, data string) {
	r := rand.New(rand.NewSource(int64(i)))
	for {
		m.Lock()
		idx := index
		index += 1
		m.Unlock()
		b := NewBoard(data, r, idx)
		b.Solve()
		if b.Solved() {
			b.Print()
			os.Exit(0)
		}
	}
}
