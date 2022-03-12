package main

import (
	"bufio"
	"fmt"
	"os"
)

func displayCount(num uint8) {
	fmt.Printf("count: %d\n", num)
	fmt.Print("(increment => i, clear => c, exit => e) >")
}

func main() {
	fmt.Print("(increment => i, clear => c, exit => e) >")
	scanner := bufio.NewScanner(os.Stdin)

	var count uint8 = 0

	for {
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "", "i", "I":
			count++
			fmt.Println("increment counter")
			displayCount(count)
			continue
		case "c", "C":
			count = 0
			fmt.Println("clear counter")
			displayCount(count)
			continue
		case "e", "E":
			fmt.Println("exit counter")
			goto L
		default:
			fmt.Println("try again")
			displayCount(count)
			continue
		}
	}
L:
}
