package main

import "fmt"

type airplane interface {
	arrive()
	depart()
	permitArrival()
}

type mediator interface {
	canArrive(airplane) bool
	notifyAboutDeparture()
}

//Mediator, control tower
type airportManager struct {
	isRunwayFree bool
	airportQueue []airplane
}

func newStationManager() *airportManager {
	return &airportManager{
		isRunwayFree: true,
	}
}

func (am *airportManager) canArrive(a airplane) bool {
	if am.isRunwayFree {
		am.isRunwayFree = false
		return true
	}
	am.airportQueue = append(am.airportQueue, a)
	return false
}

func (am *airportManager) notifyAboutDeparture() {
	if !am.isRunwayFree {
		am.isRunwayFree = true
	}
	if len(am.airportQueue) > 0 {
		nextAirplane := am.airportQueue[0]
		am.airportQueue = am.airportQueue[1:]
		nextAirplane.permitArrival()
	}
}

type plane struct {
	name     string
	mediator mediator
}

func newAirplane(name string, tower mediator) airplane {
	return &plane{
		name:     name,
		mediator: tower,
	}
}

func (p *plane) arrive() {
	if !p.mediator.canArrive(p) {
		fmt.Printf("%s, Landing blocked, waiting\n", p.name)
		return
	}
	fmt.Printf("%s: Arrived\n", p.name)
}

func (p *plane) depart() {
	fmt.Printf("%s, Taking off\n", p.name)
	p.mediator.notifyAboutDeparture()
}

func (p *plane) permitArrival() {
	fmt.Printf("%s, Landing permitted\n", p.name)
	p.arrive()
}

func main() {
	stationManager := newStationManager()

	airbus := newAirplane("airbus", stationManager)
	boeing := newAirplane("boeing", stationManager)

	airbus.arrive()
	boeing.arrive()
	airbus.depart()

}
