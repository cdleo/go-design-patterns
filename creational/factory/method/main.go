package main

import "fmt"

type dataParser interface {
	toString() string
}

type robot struct {
	name     string
	kind     string
	autonomy float32
}

func (r robot) toString() string {
	return fmt.Sprintf("Name: %s\n Kind: %s\n Autonomy: %2.2f\n",
		r.name, r.kind, r.autonomy)
}

func newRobot(name, kind string, autonomy float32) robot {
	return robot{
		name:     name,
		kind:     kind,
		autonomy: autonomy,
	}
}

func main() {
	teacherRobot := newRobot("mr. teacher robot", "teacher", 90)
	fighterRobot := newRobot("mr. fighter robot", "fighter", 70)

	fmt.Println(teacherRobot.toString())
	fmt.Println(fighterRobot.toString())

}
