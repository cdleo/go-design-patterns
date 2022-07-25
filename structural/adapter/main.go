package main

import "fmt"

type transport interface {
	navigateToDestination()
}

//struct to simulate a requirement of "car needs to go on water"
type client struct{}

func (c *client) startNavigation(trans transport) {
	fmt.Println("Client starting the navigation process")
	trans.navigateToDestination()
}

type boat struct{}

func (b *boat) navigateToDestination() {
	fmt.Println("Boat is navigating to island.")
}

type car struct{}

func (c *car) driveToDestination() {
	fmt.Println("Car is going to the destination.")
}

type carAdapter struct {
	carTransport *car
}

func (c *carAdapter) navigateToDestination() {
	fmt.Println("Adapter modify car to allow navigation")
	c.carTransport.driveToDestination()
}

func main() {
	client := &client{}
	boat := &boat{}

	client.startNavigation(boat)

	car := &car{}
	carAdapter := &carAdapter{
		carTransport: car,
	}

	client.startNavigation(carAdapter)
}
