package main

import (
	"log"

	subject_observer "github.com/observer-design-pattern/subject-observer"
)

func main() {
	order := &subject_observer.Order{ID: "12345"}
	restaurant := &subject_observer.RestaurantObserver{Name: "Pizza Place"}
	customer := &subject_observer.CustomerObserver{Name: "John Doe"}
	delivery := &subject_observer.DeliveryObserver{Name: "Fast Delivery"}
	callCenter := &subject_observer.CallCenterObserver{Name: "Support Center"}

	order.AddObserver(restaurant)
	order.AddObserver(customer)
	order.AddObserver(delivery)	
	order.AddObserver(callCenter)

	order.NotifyObservers()

	log.Println("----------------------------------------------")

	order.RemoveObserver(customer)

	order.NotifyObservers()
}