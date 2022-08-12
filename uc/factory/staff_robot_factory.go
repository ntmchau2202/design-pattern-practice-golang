package factory

import (
	"design-pattern/entities/robot"
	"design-pattern/entities/robot/components/hand"
	"design-pattern/uc/builder"
)

type StaffRobotFactory struct {
	robot *robot.StaffRobot
	RobotFactory
}

func (f *StaffRobotFactory) Init() {
	f.robot = &robot.StaffRobot{}
	f.typeBuilder = &builder.StaffRobotBuilder{}
}

func (f *StaffRobotFactory) CreateNaiveRobot() {
	f.Init()
	f.levelBuilder = &builder.NaiveRobotBuilder{}

	f.Build(f.robot)
}

func (f *StaffRobotFactory) CreateIntermediateRobot() {
	f.Init()
	f.levelBuilder = &builder.IntermediateRobotBuilder{}

	f.Build(f.robot)

	// reset hands
	f.robot.SetHands(hand.SkillfulHand{}, hand.SkillfulHand{})

}

func (f *StaffRobotFactory) CreateAdvancedRobot() {
	f.Init()
	f.levelBuilder = &builder.AdvancedRobotBuilder{}

	f.Build(f.robot)
	// reset hands
	f.robot.SetHands(hand.SkillfulHand{}, hand.SkillfulHand{})
}
