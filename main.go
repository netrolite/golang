package main

import (
	"bufio"
	"fmt"
	"github.com/netrolite/golang/utils"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.SetFlags(log.Ltime)
	rd := bufio.NewReader(os.Stdin)
	order := utils.Order{Items: utils.OrderItems{}}

	fmt.Print("Your name: ")
	order.Customer.Name = utils.ReadString(rd)

	fmt.Print("Your age: ")
	order.Customer.Age = utils.ReadInt(rd)

	for true {
		itemKeys := utils.ItemKeys()
		if len(order.Items) >= 1 {
			fmt.Println("Items in your order:")
			for itemName, itemQty := range order.Items {
				fmt.Printf("%vx %v ($%v each)\n", itemQty, itemName, utils.Items[itemName])
			}
		}

		if len(order.Items) < 1 {
			fmt.Printf("Select an item that you would like to add to your order (1-%v):\n", len(utils.Items))
		} else {
			fmt.Printf("Select an item that you would like to add to your order (1-%v) or press enter to continue:\n", len(utils.Items))
		}
		for i, item := range itemKeys {
			fmt.Printf("(%v) %v: $%v\n", i+1, item, utils.Items[item])
		}

		selectedItemNumStr, err := rd.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(order.Items) < 1 && selectedItemNumStr == "\n" {
			fmt.Println("You haven't selected any items!")
			continue
		} else if selectedItemNumStr == "\n" {
			break
		}
		selectedItemNumStr = strings.TrimSpace(selectedItemNumStr)

		selectedItemNum, err := strconv.Atoi(selectedItemNumStr)
		if err != nil {
			log.Fatal(err)
		}
		selectedItemNum--
		if selectedItemNum < 0 || selectedItemNum > len(utils.Items) {
			log.Fatal("Item number is out of range")
		}

		selectedItemKey := itemKeys[selectedItemNum]
		fmt.Printf("How many of %v do you need? ", selectedItemKey)
		selectedItemQty := utils.ReadInt(rd)
		if selectedItemQty < 1 {
			log.Fatal("Item quantity cannot be less than 1")
		}
		order.Items[selectedItemKey] = order.Items[selectedItemKey] + selectedItemQty
	}

	order.PrintReceipt()
}
