package main

import (
	"fmt"
)

func main() {
	var cookiesShop Shop = Shop{100, 5, 0}
	var queue []Queue = []Queue{{"олег", 0, 10}, {"вова", 0, 15}, {"акула", 0, 0}}
	queue[1] = queue[1].SetHaveBuy(1)

	queue = queue[1:]

	fmt.Println(queue, cookiesShop)
}

type Shop struct {
	QuantitySale int
	Price        int
	Money        int
}

type Queue struct {
	Name    string
	Money   int
	HaveBuy int
}

func (q Queue) SetHaveBuy(i int) Queue {
	q.HaveBuy = i
	return q
}
