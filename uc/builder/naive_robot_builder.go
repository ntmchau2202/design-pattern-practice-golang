package builder

import (
	"design-pattern/entities/robot"
	"design-pattern/entities/robot/components/hand"
	"design-pattern/entities/robot/components/wheel"
)

type NaiveRobotBuilder struct {
	RobotBuilder
}

func (b *NaiveRobotBuilder) BuildHands(r robot.WorkableRobot) {
	r.SetHands(hand.NormalHand{}, hand.NormalHand{})
}

func (b *NaiveRobotBuilder) BuildWheels(r robot.WorkableRobot) {
	r.SetWheels(wheel.StableWheel{}, wheel.StableWheel{}, wheel.StableWheel{}, wheel.StableWheel{})
}

func (b *NaiveRobotBuilder) BuildCabins(r robot.WorkableRobot) {

}
