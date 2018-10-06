
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	x, y, NT := 50000, 10000, 5000

	sq := func() {
		defer wg.Done()
		v := 0
		for i:=0; i<x; i++ {
			for j:=0; j<y; j++ {
				v = j * j * j
			}
		}
		fmt.Println(v)
	}

	for i:=0; i<NT; i++ {
		go sq()
		wg.Add(1)
	}

	wg.Wait()
}
