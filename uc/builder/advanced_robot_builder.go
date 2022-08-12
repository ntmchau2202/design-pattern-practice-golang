package builder

import (
	"design-pattern/entities/robot"
	"design-pattern/entities/robot/components/cabin"
	"design-pattern/entities/robot/components/wheel"
)

type AdvancedRobotBuilder struct {
	RobotBuilder
}

func (b *AdvancedRobotBuilder) BuildHands(r robot.WorkableRobot) {

}

func (b *AdvancedRobotBuilder) BuildWheels(r robot.WorkableRobot) {
	r.SetWheels(wheel.StableWheel{}, wheel.StableWheel{}, wheel.QuickWheel{}, wheel.QuickWheel{})
}

func (b *AdvancedRobotBuilder) BuildCabins(r robot.WorkableRobot) {
	r.SetCabins(cabin.MainCabin{}, nil)

}
