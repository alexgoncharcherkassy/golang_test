package main

import (
	"awesomeProject/src/ekreative.com/machine"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	vehicleMachine := machine.MakeMachine(4, 5)

	fmt.Println("Enter: exit - close app; 1 - add product; 2 - make an order; 3 - purchase")
	fmt.Print("Choose operation: ")
	scanner := bufio.NewScanner(os.Stdin)

	exit := false
	for scanner.Scan() {
		switch scanner.Text() {
		case "exit":
			exit = true
			break
		case "1":
			enteredData := ParseAdding()

			for key, val := range enteredData {
				bucket := strings.Split(val, ",")

				for _, item := range bucket {
					if item != "" {
						parsedItem := strings.Split(item, "->")

						if len(parsedItem) == 2 {
							if parsedItem[0] != "" && parsedItem[1] != "" {
								name, err1 := strconv.Atoi(parsedItem[0])
								price, err2 := strconv.Atoi(parsedItem[1])

								if err1 != nil || err2 != nil {
									fmt.Println(err1)
									fmt.Println(err2)
								} else {
									AddProduct(vehicleMachine, name, key+1, price)
								}
							}
						}
					}
				}
			}

			PrintMachine(vehicleMachine)
			break
		case "2":
			enteredData := ParseGetting()
			var products []int

			for _, val := range enteredData {
				i1, err := strconv.Atoi(val)
				if err != nil {
					fmt.Println(err)
				} else {
					products = append(products, i1)
				}
			}

			GetProducts(vehicleMachine, products, true)
			break
		case "3":
			GetProducts(vehicleMachine, []int{}, false)
			break
		}

		if exit == true {
			break
		}

		fmt.Print("Choose operation: ")
	}
}

func AddProduct(m *machine.Machine, nameProduct int, numberOfBucket int, price int) {
	err := m.AddProduct(nameProduct, numberOfBucket, price)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetProducts(m *machine.Machine, products []int, preSale bool) {
	foundProducts, sum, err := m.GetProducts(products, preSale)

	if err != nil {
		fmt.Println(err)
		return
	}

	PrintSum(sum)
	if preSale == false {
		PrintFoundProducts(foundProducts)
		PrintMachine(m)
	}
}

func ParseAdding() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("fill buckets: ")

	text, err := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	return strings.Split(text, ";")
}

func ParseGetting() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter order: ")

	text, err := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	if err != nil {
		fmt.Println(err)
	}

	return strings.Split(text, ",")
}

func PrintMachine(m *machine.Machine) {
	fmt.Println("-------------------------")

	for i := 1; i <= len(m.Buckets()); i++ {
		fmt.Print("[ ")
		for _, val := range m.BucketProducts(i) {
			fmt.Printf("%v ", val.Name())
		}
		fmt.Print("] ")
		fmt.Printf("(%v)  ", m.GetCurrentNumberOfProducts(i))
	}
	fmt.Printf("  Total number of products in machine: %v", m.GetCurrentNumberOfProducts(-1))
	fmt.Println()
}

func PrintFoundProducts(products []int) {
	fmt.Print("Items in order getting: ")

	for _, val := range products {
		fmt.Printf(" %v ", val)
	}

	fmt.Println()
}

func PrintSum(sum int) {
	fmt.Printf("Total order sum: %v\n", sum)
}
