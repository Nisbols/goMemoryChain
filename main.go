// A Pattern matching game
//
package main

import "fmt"
import "math/rand"
import "time"
import "os"
import "log"

const maxNumber = 4
const minNumber = 1

func main() {
	ranNumber()
	file, err := os.Create("doc")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "%d", ranNumber())
	fmt.Fprintf(file, "%d", ranNumber())
}

// A random number generator with a new seed ever nano second
// The range can be changed
func ranNumber() int {
	rand.Seed(time.Now().UnixNano())
	n := 0
	n = rand.Intn(maxNumber-minNumber) + minNumber
	return n
}
