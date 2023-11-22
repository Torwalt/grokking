package linkedlist

import (
	"fmt"
	"strings"
)

type Order struct {
	Item  string
	Price int
}

func (o Order) String() string {
	return fmt.Sprintf("%v for %v€", o.Item, o.Price)
}

type OrderItem struct {
	Order Order
	Next  *OrderItem
}

type OrderList struct {
	Head *OrderItem
}

func (ol *OrderList) String() string {
	out := []string{}

	o := ol.Head
	for o.Next != nil {
		out = append(out, o.Order.String())
		o = o.Next
	}

	// Append the last one.
	out = append(out, o.Order.String())

	return strings.Join(out, "->")
}

func (ol *OrderList) Insert(order Order) {
	item := &OrderItem{
		Order: order,
		Next:  nil,
	}

	if ol.Head == nil {
		ol.Head = item

		return
	}

	n := ol.Head
	for n.Next != nil {
		n = n.Next
	}

	n.Next = item
}

func NewOrderList(orders ...Order) *OrderList {
	if len(orders) == 0 {
		return &OrderList{}
	}

	out := &OrderList{
		Head: nil,
	}

	for _, order := range orders {
		out.Insert(order)
	}

	return out
}
