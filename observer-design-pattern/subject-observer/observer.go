package subject_observer

import "log"

type Observer interface {
	SendNotification(order Order)
	GetName() string
}

type RestaurantObserver struct {
	Name string
}

func (r *RestaurantObserver) SendNotification(order Order) {
	log.Printf("Restaurant notified for order ID: %s\n", order.ID)
}

func (r *RestaurantObserver) GetName() string {
	return r.Name
}

type CustomerObserver struct {
	Name string
}

func (c *CustomerObserver) SendNotification(order Order) {
	log.Printf("Customer notified for order ID: %s\n", order.ID)
}

func (c *CustomerObserver) GetName() string {
	return c.Name
}

type DeliveryObserver struct {
	Name string
}

func (d *DeliveryObserver) SendNotification(order Order) {
	log.Printf("Delivery person notified for order ID: %s\n", order.ID)
}

func (d *DeliveryObserver) GetName() string {
	return d.Name
}

type CallCenterObserver struct {
	Name string
}

func (c *CallCenterObserver) SendNotification(order Order) {
	log.Printf("Call center notified for order ID: %s\n", order.ID)
}

func (c *CallCenterObserver) GetName() string {
	return c.Name
}