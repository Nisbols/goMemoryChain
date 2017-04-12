// A Pattern matching game
//
package main

import "math/rand"
import "time"

import "fmt"

const maxNumber = 4
const minNumber = 1

func main() {
	c := make([]int, 0)
	numberLoop(c)
	fmt.Printf("exited loop")
}

// A random number generator with a new seed ever nano second
// The range can be changed
func ranNumber() int {
	rand.Seed(time.Now().UnixNano())
	n := 0
	n = rand.Intn(maxNumber-minNumber) + minNumber
	return n
}

func numberLoop(c []int) {
	for l := 1; l == 1; {
		u := 0
		n := len(c)
		c = append(c, ranNumber())
		//fmt.Println(n)
		//fmt.Println(c)
		for i := 0; i <= n; i++ {
			fmt.Printf("%d", c[i])
			time.Sleep(1)
		}
		for i := 0; (i <= n) && (l == 1); i++ {
			fmt.Printf("\n")
			fmt.Scanf("%d", &u)
			if u != c[i] {
				fmt.Printf("bad\n")
				l = 0
			}
		}
	}
}
