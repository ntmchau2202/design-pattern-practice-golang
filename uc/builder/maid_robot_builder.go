package builder

import (
	"design-pattern/entities/robot"
	"design-pattern/entities/robot/components/cabin"
)

type MaidRobotBuilder struct {
	RobotBuilder
}

func (b *MaidRobotBuilder) BuildHands(r robot.WorkableRobot) {
}

func (b *MaidRobotBuilder) BuildWheels(r robot.WorkableRobot) {
}

func (b *MaidRobotBuilder) BuildCabins(r robot.WorkableRobot) {
	r.SetCabins(cabin.MainCabin{}, nil)
}
