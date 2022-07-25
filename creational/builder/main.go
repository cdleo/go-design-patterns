package main

import "fmt"

type robot struct {
	rType     string
	autonomy  float32
	hasGuitar bool
	hasGun    bool
	hasHelm   bool
	hasHat    bool
	hasSword  bool
}

type robotBuilder struct {
	robot robot
}

func newRobotBuilder() *robotBuilder {
	return &robotBuilder{
		robot: robot{},
	}
}

func (r *robotBuilder) setAutonomy(autonomy float32) *robotBuilder {
	r.robot.autonomy = autonomy
	return r
}

func (r *robotBuilder) setType(rType string) *robotBuilder {
	r.robot.rType = rType
	return r
}

func (r *robotBuilder) withGuitar() *robotBuilder {
	r.robot.hasGuitar = true
	return r
}

func (r *robotBuilder) withGun() *robotBuilder {
	r.robot.hasGun = true
	return r
}

func (r *robotBuilder) withHelm() *robotBuilder {
	r.robot.hasHelm = true
	return r
}

func (r *robotBuilder) withHat() *robotBuilder {
	r.robot.hasHat = true
	return r
}

func (r *robotBuilder) withSword() *robotBuilder {
	r.robot.hasSword = true
	return r
}

func (r *robotBuilder) build() robot {
	return r.robot
}

func main() {

	robotFigther := newRobotBuilder().
		setAutonomy(100.00).
		setType("figther").
		withGun().
		withHelm().
		withSword().
		build()

	robotSinger := newRobotBuilder().
		setAutonomy(100.00).
		setType("singer").
		withGuitar().
		withHat().
		build()

	printRobotInfo(robotFigther)
	fmt.Println("----------------------------")
	printRobotInfo(robotSinger)
}

func printRobotInfo(robot robot) {

	fmt.Printf("Type: %s\n Autonomy: %f\n HasGun: %t\n HasGuitar: %t\n HasHelm: %t\n HasHat: %t\n HasSword: %t\n",
		robot.rType, robot.autonomy, robot.hasGun, robot.hasGuitar, robot.hasHelm, robot.hasHat, robot.hasSword)
}
