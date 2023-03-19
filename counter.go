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
		case "increment", "", "i", "I":
			count++
			fmt.Println("increment count")
			displayCount(count)
			continue
		case "decrement", "d", "D":
			if count > 0 {
				count--
				fmt.Println("decrement count")
			} else {
				fmt.Println("count is less than or equal to 0")
			}
			displayCount(count)
			continue
		case "clear" ,"c", "C":
			count = 0
			fmt.Println("clear count")
			displayCount(count)
			continue
		case "exit", "e", "E":
			fmt.Println("exit counter")
			break Loop
		default:
			fmt.Println("counter: command not found")
			displayCount(count)
			continue
		}
	}
}
