package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	count := 0
	m := map[int]bool{}
	m[0] = true
	m[1] = true
	for {
		var n int
		for {
			n = rand.Int() % 203
			if m[n] {
				continue
			} else {
				break
			}
		}

		m[n] = true
		count++
		if count == 201 {
			return
		}
		fmt.Println(n)
	}

}
