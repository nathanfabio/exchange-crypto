package main

import (
	"fmt"
	"testing"
)


func TestLimit(t *testing.T) {
	l := NewLimit(100.0)
	buyOrder := NewOrder(true, 1.0)

	l.AddOrder(buyOrder)

	fmt.Println(l)
}

func TestOrderbook(t *testing.T) {

}