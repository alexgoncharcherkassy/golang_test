package main

import (
	"awesomeProject/src/ekreative.com/machine"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func parseBuckets(vendingMachine *machine.Machine, str string) {
	if str == "" {
		return
	}

	enteredData := strings.Split(str, ";")

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
							AddProduct(vendingMachine, name, key+1, price)
						}
					}
				}
			}
		}
	}
}

func parseItems(str string) []int {
	if str == "" {
		return []int{}
	}
	strs := strings.Split(str, ",")
	var items []int

	for _, val := range strs {
		i1, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println(err)
		} else {
			items = append(items, i1)
		}
	}

	return items
}

func Test_GetProducts(t *testing.T) {
	type args struct {
		buckets string
		order   []int
	}
	tests := []struct {
		name string
		args args
		want bool
		sum  int
	}{
		{name: "1", args: args{"1->1,2->2,3->3,5->5,5->5;2->2,5->5,4->4,3->3,1->1;3->3,5->5,4->4,1->1,1->1;5->5,1->1,1->1,1->1,1->1", parseItems("1,2,3,4,5")}, want: true, sum: 15},
		{name: "2", args: args{"1->1,1->1,3->3;1->1,1->1,2->2", parseItems("1,2,1")}, want: true, sum: 4},
		{name: "3", args: args{"1->1", parseItems("1")}, want: true, sum: 1},
		{name: "4", args: args{"1->1", parseItems("2")}, want: false, sum: 0},
		{name: "5", args: args{"1->1,2->2", parseItems("2")}, want: false, sum: 0},
		{name: "6", args: args{"1->1,2->2", parseItems("1,2")}, want: true, sum: 3},
		{name: "7", args: args{"1->1,2->2", parseItems("2,1")}, want: true, sum: 3},
		{name: "8", args: args{"1->1,2->2;3->3,4->4", parseItems("3")}, want: true, sum: 3},
		{name: "9", args: args{"1->1,2->2;3->3,4->4", parseItems("4")}, want: false, sum: 0},
		{name: "10", args: args{"1->1,2->2;3->3,4->4", parseItems("3,4")}, want: true, sum: 7},
		{name: "11", args: args{"1->1,2->2;3->3,4->4", parseItems("3,4,2")}, want: false, sum: 0},
		{name: "12", args: args{"1->1,2->2;3->3,4->4", parseItems("3,4,2,1")}, want: true, sum: 10},
		{name: "13", args: args{"1->1;1->1,2->2", parseItems("1,2")}, want: true, sum: 3},
		{name: "14", args: args{"1->1,2->2,3->3,4->4;1->1,2->2;1->1,2->2,3->3,5->5;1->1,2->2", parseItems("1,1,2,2,3,3,4,5")}, want: true, sum: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vendingMachine := machine.MakeMachine(4, 5)
			parseBuckets(vendingMachine, tt.args.buckets)
			if got, sum, err := vendingMachine.GetProducts(tt.args.order, true); err != nil && tt.want || err == nil && !tt.want {
				t.Errorf("got = %v, sum = %v, want %v", got, sum, tt.want)
			}
			if got, sum, err := vendingMachine.GetProducts(tt.args.order, false); err != nil && len(got) > 0 && sum != tt.sum || err == nil && len(got) != len(tt.args.order) && sum != tt.sum {
				t.Errorf("got = %v, sum = %v, want %v", got, sum, tt.want)
			}
		})
	}
}
