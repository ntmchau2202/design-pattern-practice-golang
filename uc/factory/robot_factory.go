package factory

import (
	"design-pattern/entities/robot"
	"design-pattern/uc/builder"
)

type WorkableRobotFactory interface {
	Init()
	CreateNaiveRobot()
	CreateIntermediateRobot()
	CreateAdvancedRobot()
}

type RobotFactory struct {
	WorkableRobotFactory
	typeBuilder  builder.RobotBuilder
	levelBuilder builder.RobotBuilder
}

func (f *RobotFactory) Build(r robot.WorkableRobot) {
	f.levelBuilder.BuildHands(r)
	f.levelBuilder.BuildWheels(r)
	f.levelBuilder.BuildCabins(r)

	f.typeBuilder.BuildCabins(r)
	f.typeBuilder.BuildWheels(r)
	f.typeBuilder.BuildCabins(r)
}
