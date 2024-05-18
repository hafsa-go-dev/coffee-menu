package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
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

	input, err := in.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		fmt.Println(err)
	}
loop:
	for {
		switch input = strings.TrimSpace(input); input {
		case "1":
			for _, item := range menu {
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 22))
				for size, price := range item.price {
					fmt.Printf("%10s: %10.2f\n", size, price)
				}
			}
			break loop
		case "2":
			fmt.Println("Please enter the name of the item you would like to add:")
			input, err := in.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}
			menu = append(menu, menuItem{name: input, price: make(map[string]float64)})
			fmt.Println("Please enter the price of the item for a size small:")
			input, err = in.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}
			input = strings.TrimSpace(input)
			numInput, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("That is not a valid number.")
			} else {
				menu[len(menu)-1].price["small"] = numInput
				fmt.Println("Please enter the price of the item for a size medium:")
				input, err = in.ReadString('\n')
				if err != nil {
					fmt.Println(err)
				}
				numInput, err := strconv.ParseFloat(input, 64)
				if err != nil {
					fmt.Println("That is not a valid number.")
				} else {
					menu[len(menu)-1].price["medium"] = numInput
					fmt.Println("Please enter the price of the item for a size large:")
					input, err = in.ReadString('\n')
					if err != nil {
						fmt.Println(err)
					}
					numInput, err := strconv.ParseFloat(input, 64)
					if err != nil {
						fmt.Println("That is not a valid number.")
					} else {
						menu[len(menu)-1].price["large"] = numInput
					}
				}
			}
			break loop
		case "3":
			fmt.Println("Please enter the name of the item you would like to remove:")
			input, err := in.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}

			for i, v := range menu {
				if v.name == input {
					slices.Delete(menu, i, i+1)
				}
			}
			break loop
		case "Q", "q":
			fmt.Println("Exiting the program...")
			break loop
		}
	}
}
