package linkedlist_test

import (
	"grokking/exercises/linkedlist"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewOrderList(t *testing.T) {
	orders := []linkedlist.Order{
		{Item: "Pizza", Price: 5},
		{Item: "Spaghetti", Price: 7},
		{Item: "Scampi", Price: 6},
		{Item: "Mozarella", Price: 9},
		{Item: "Focaccio", Price: 4},
	}

	ll := linkedlist.NewOrderList(orders...)
	require.NotEmpty(t, ll)

	stringOrders := ll.String()
	require.NotEmpty(t, stringOrders)

	pizza := ll.Pop()
	require.Equal(t, "Pizza", pizza.Item)

	spaghetti := ll.Pop()
	require.Equal(t, "Spaghetti", spaghetti.Item)

	leftOvers := ll.String()
	require.NotEmpty(t, leftOvers)
	require.Less(t, len(leftOvers), len(stringOrders))
}
