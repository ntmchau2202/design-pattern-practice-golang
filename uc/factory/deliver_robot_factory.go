package factory

import (
	"design-pattern/entities/robot"
	"design-pattern/entities/robot/components/cabin"
	"design-pattern/uc/builder"
)

type DeliverRobotFactory struct {
	robot *robot.DeliverRobot
	RobotFactory
}

func (f *DeliverRobotFactory) Init() {
	f.robot = &robot.DeliverRobot{}
	f.typeBuilder = &builder.DeliverRobotBuilder{}
}

func (f *DeliverRobotFactory) CreateNaiveRobot() *robot.DeliverRobot {
	f.Init()
	f.levelBuilder = &builder.NaiveRobotBuilder{}

	f.Build(f.robot)
	return f.robot
}

func (f *DeliverRobotFactory) CreateIntermediateRobot() *robot.DeliverRobot {
	f.Init()
	f.levelBuilder = &builder.IntermediateRobotBuilder{}

	f.Build(f.robot)

	// add 1 extra cabin
	f.robot.SetCabins(nil, cabin.SubCabin{})

	return f.robot
}

func (f *DeliverRobotFactory) CreateAdvancedRobot() *robot.DeliverRobot {
	f.Init()
	f.levelBuilder = &builder.AdvancedRobotBuilder{}

	f.Build(f.robot)

	// add 1 extra cabin
	f.robot.SetCabins(nil, cabin.SubCabin{})

	return f.robot
}
