package main

import (
	"bufio"
	"fmt"
	"os"
)

func displayCount(num uint8) {
	fmt.Printf("count: %d\n", num)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var count uint8 = 0

	Loop: for {
		fmt.Print("Enter a value (type 'exit' to quit): ")
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
		case "exit", "e", "E":
			fmt.Println("exit counter")
			break Loop
		default:
			fmt.Println("try again")
			displayCount(count)
			continue
		}
	}
}
