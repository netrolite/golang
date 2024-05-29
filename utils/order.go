package utils

import (
	"fmt"
	"sort"
)

type Order struct {
	Customer struct {
		Name string
		Age  int
	}
	Items OrderItems
	Tip   float32
}

type OrderItems map[string]int

var Items = map[string]float32{
	"banana":      30,
	"apple":       10,
	"apple juice": 15,
	"pineapple":   60,
	"coconut":     40,
}

func (o *Order) PrintReceipt() {
	var total float32 = 0

	itemsList := ""
	for item, qty := range o.Items {
		price := Items[item]
		total += price * float32(qty)
		itemsList += fmt.Sprintf("%vx %v ($%v each)\n", qty, item, price)
	}

	fmt.Printf(`Receipt:
Customer name: %v
Customer age: %v
%vTotal: $%v
See you next time!
`, o.Customer.Name, o.Customer.Age, itemsList, total)
}

func ItemKeys() []string {
	itemKeys := make([]string, len(Items))

	i := 0
	for k := range Items {
		itemKeys[i] = k
		i++
	}
	sort.Strings(itemKeys)

	return itemKeys
}
