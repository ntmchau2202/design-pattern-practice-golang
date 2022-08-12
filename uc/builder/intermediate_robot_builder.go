package builder

import (
	"design-pattern/entities/robot"
	"design-pattern/entities/robot/components/cabin"
	"design-pattern/entities/robot/components/wheel"
)

type IntermediateRobotBuilder struct {
	RobotBuilder
}

func (b *IntermediateRobotBuilder) BuildHands(r robot.WorkableRobot) {

}

func (b *IntermediateRobotBuilder) BuildWheels(r robot.WorkableRobot) {
	r.SetWheels(wheel.StableWheel{}, wheel.StableWheel{}, wheel.StableWheel{}, wheel.StableWheel{})
}

func (b *IntermediateRobotBuilder) BuildCabins(r robot.WorkableRobot) {
	r.SetCabins(cabin.MainCabin{}, nil)

}
