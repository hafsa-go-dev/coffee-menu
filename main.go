package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type menuItem struct {
		name  string
		price map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", price: map[string]float64{
			"small":  3.50,
			"medium": 4.50,
			"large":  5.50,
		}},
		{name: "Tea", price: map[string]float64{
			"small":  2.50,
			"medium": 3.50,
			"large":  4.50,
		}},
	}

	in := bufio.NewReader(os.Stdin)

	fmt.Println("Please make a selection:")
	fmt.Println("[1] Print menu")
	fmt.Println("[2] Add an item")
	fmt.Println("[3] Remove an item")
	fmt.Println("[Q] Quit the program")

	choice, _ := in.ReadString('\n')

	switch strings.TrimSpace(choice) {
	case "1":
		for _, item := range menu {
			fmt.Println(item.name)
			fmt.Println(strings.Repeat("-", 22))
			for size, price := range item.price {
				fmt.Printf("%10s: %10.2f\n", size, price)
			}
		}
	}
}
