package factory

import (
	"design-pattern/entities/robot"
	"design-pattern/entities/robot/components/cabin"
	"design-pattern/entities/robot/components/hand"
	"design-pattern/uc/builder"
)

type MaidRobotFactory struct {
	robot *robot.MaidRobot
	RobotFactory
}

func (f *MaidRobotFactory) Init() {
	f.robot = &robot.MaidRobot{}
	f.typeBuilder = &builder.MaidRobotBuilder{}
}

func (f *MaidRobotFactory) CreateNaiveRobot() *robot.MaidRobot {
	f.Init()
	f.levelBuilder = &builder.NaiveRobotBuilder{}

	f.Build(f.robot)
	return f.robot
}

func (f *MaidRobotFactory) CreateIntermediateRobot() *robot.MaidRobot {
	f.Init()
	f.levelBuilder = &builder.IntermediateRobotBuilder{}

	f.Build(f.robot)

	// add 1 cabin
	f.robot.SetCabins(nil, cabin.SubCabin{})

	// reset hands
	f.robot.SetHands(hand.SkillfulHand{}, hand.SkillfulHand{})
	return f.robot
}

func (f *MaidRobotFactory) CreateAdvancedRobot() *robot.MaidRobot {
	f.Init()
	f.levelBuilder = &builder.AdvancedRobotBuilder{}

	f.Build(f.robot)

	// add 1 cabin
	f.robot.SetCabins(nil, cabin.SubCabin{})

	// reset hands
	f.robot.SetHands(hand.SkillfulHand{}, hand.SkillfulHand{})
	return f.robot
}
