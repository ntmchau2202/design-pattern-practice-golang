package builder

import (
	"design-pattern/entities/robot"
	"design-pattern/entities/robot/components/cabin"
	"design-pattern/entities/robot/components/hand"
)

type DeliverRobotBuilder struct {
	RobotBuilder
}

func (b *DeliverRobotBuilder) BuildHands(r robot.WorkableRobot) {
	r.SetHands(hand.NormalHand{}, hand.NormalHand{})
}

func (b *DeliverRobotBuilder) BuildWheels(r robot.WorkableRobot) {
}

func (b *DeliverRobotBuilder) BuildCabins(r robot.WorkableRobot) {
	r.SetCabins(cabin.MainCabin{}, nil)
}
