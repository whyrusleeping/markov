package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/whyrusleeping/markov"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("must pass in corpus file")
		return
	}

	infi := os.Args[1]
	fi, err := os.Open(infi)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	n := new(markov.Node)
	scan := bufio.NewScanner(fi)
	for scan.Scan() {
		n.InsertPhrase(scan.Text())
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		fmt.Println(n.GeneratePhrase())
	}
}
