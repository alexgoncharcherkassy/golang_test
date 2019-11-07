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
	vendingMachine := machine.MakeMachine(4, 5)

	fmt.Println("Enter: exit - for close app; 1 - for add product; 2 - for get products")
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

				for _, name := range bucket {
					if name != "" {
						i1, err := strconv.Atoi(name)
						if err != nil {
							fmt.Println(err)
						} else {
							AddProduct(vendingMachine, i1, key+1)
						}
					}
				}
			}

			PrintMachine(vendingMachine)
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

			GetProducts(vendingMachine, products)
			break
		}

		if exit == true {
			break
		}

		fmt.Print("Choose operation: ")
	}
}

func AddProduct(m machine.Machine, nameProduct int, numberOfBucket int) {
	err := m.AddProduct(nameProduct, numberOfBucket)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetProducts(m machine.Machine, products []int) {
	foundProducts, err := m.GetProducts(products)

	if err != nil {
		fmt.Println(err)
		return
	}

	if len(foundProducts) == 0 {
		fmt.Println("impossible")
		return
	}

	PrintFoundProducts(foundProducts)
	PrintMachine(m)
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

func PrintMachine(m machine.Machine) {
	fmt.Println("-------------------------")

	for i := 1; i <= len(m.Buckets()); i++ {
		fmt.Print("[ ")
		for _, val := range m.Buckets()[i].Elements() {
			fmt.Printf("%v ", val.Name())
		}
		fmt.Print("] ")
		fmt.Printf("(%v)  ", len(m.Buckets()[i].Elements()))
	}
	fmt.Printf("  Total number of products: %v", m.GetCurrentNumberOfProducts())
	fmt.Println()
}

func PrintFoundProducts(products []int) {
	for _, val := range products {
		fmt.Printf(" %v ", val)
	}
	fmt.Println()
}
