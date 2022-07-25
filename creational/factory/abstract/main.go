package main

import "fmt"

const (
	teacherType string = "teacherRobot"
	fighterType string = "fighterRobot"
)

type imRobot interface {
	whoAmI() string
}

type robot struct {
	name     string
	rType    string
	autonomy float32
}

func (r robot) whoAmI() string {
	return fmt.Sprintf("Name: %s\n rType: %s\n Autonomy: %2.2f",
		r.name, r.rType, r.autonomy)
}

type teacherRobot struct {
	robot
}

func newTeacherRobot() imRobot {
	return teacherRobot{
		robot: robot{
			name:     "Mr. Teacher",
			rType:    "teacher",
			autonomy: 100.00,
		},
	}
}

type fighterRobot struct {
	robot
}

func newFighterRobot() imRobot {
	return fighterRobot{
		robot: robot{
			name:     "Mr. Fighter",
			rType:    "fighter",
			autonomy: 70.00,
		},
	}
}

func createRobot(rType string) (imRobot, error) {
	if rType == teacherType {
		return newTeacherRobot(), nil
	}
	if rType == fighterType {
		return newFighterRobot(), nil
	}
	return nil, fmt.Errorf("Could not create a robot of type %s", rType)
}

func main() {
	robotA, err := createRobot(teacherType)
	robotB, err := createRobot(fighterType)
	_, err = createRobot("invalid")

	fmt.Println(robotA.whoAmI())
	fmt.Println("------------------")
	fmt.Println(robotB.whoAmI())
	fmt.Println("------------------")
	fmt.Println("WrongRobot err: ", err.Error())

}
