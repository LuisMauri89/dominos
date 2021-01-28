package main

import (
	"fmt"
	"testing"
)

type Delivery struct {
	ID          string
	OrderID     string
	Status      string
	To          string
	FinalPrice  int
	Address     string
	Description string
}

var Del []Delivery

func main() {
	Del = append(Del, Delivery{ID: "123", OrderID: "1", Status: "Declined", To: "Juan Gabriel", FinalPrice: 1999, Address: "Santiago", Description: "Pizza"})

}

func ObtainDelivery(t *testing.T) {
	fmt.Println("ID")
	fmt.Println("OrderID")
	fmt.Println("Status")
	fmt.Println("To")
	fmt.Println("Price")
	fmt.Println("Address")
	fmt.Println("Description")

}
