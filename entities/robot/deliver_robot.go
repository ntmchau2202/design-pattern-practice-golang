package robot

import "fmt"

type DeliverRobot struct {
	Robot
	ClonableRobot
}

func (r *DeliverRobot) Greet() {
	fmt.Println("Deliver robot say hi!")
}

func (r *DeliverRobot) Bye() {
	fmt.Println("Deliver robot say bye!")
}

func (r *DeliverRobot) DoMainTask() {
	fmt.Println("Deliver robot work hard!")
}

func (r *DeliverRobot) DoSubTask() {
	fmt.Println("Deliver robot play hard!")
}

func (r *DeliverRobot) Clone() WorkableRobot {
	clone := &DeliverRobot{}
	clone.SetCabins(r.Parts.RobotCabin[0], r.Parts.RobotCabin[1])
	clone.SetHands(r.Parts.RobotHand[0], r.Parts.RobotHand[1])
	clone.SetWheels(r.Parts.RobotWheel...)
	return clone
}
