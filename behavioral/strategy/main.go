package main

import "fmt"

type outputter interface {
	sendImage()
	getDestination() string
}

type printer struct {
	destination string
}

func (p printer) sendImage() {
	fmt.Println("Printing image")
}

func (p printer) getDestination() string {
	return p.destination
}

type display struct {
	destination string
}

func (d display) sendImage() {
	fmt.Println("Drawing image in screen")
}

func (d display) getDestination() string {
	return d.destination
}

func draw(output outputter) {
	fmt.Println("Preparing destination: ", output.getDestination())
	output.sendImage()
}

func main() {
	p := printer{destination: "Printer machine"}
	d := display{destination: "Graphic card"}

	draw(p)
	fmt.Println("-------------------------------------")
	draw(d)
}
