package prototype

import (
	"design-pattern/entities/robot"
	"design-pattern/uc/factory"
)

type RobotCategory string

const (
	DELIVER_NAIVE        RobotCategory = "deliver_naive"
	DELIVER_INTERMEDIATE RobotCategory = "deliver_intermediate"
	DELIVER_ADVANCED     RobotCategory = "deliver_advanced"

	MAID_NAIVE        RobotCategory = "maid_naive"
	MAID_INTERMEDIATE RobotCategory = "maid_intermediate"
	MAID_ADVANCED     RobotCategory = "maid_advanced"

	STAFF_NAIVE        RobotCategory = "staff_naive"
	STAFF_INTERMEDIATE RobotCategory = "staff_intermediate"
	STAFF_ADVANCED     RobotCategory = "staff_advanced"
)

type robotRegistry map[RobotCategory]robot.ClonableRobot

var RobotRegistry robotRegistry

func InitRobotRgistry() {
	RobotRegistry = make(map[RobotCategory]robot.ClonableRobot)
	deliverFactory := factory.DeliverRobotFactory{}

	RobotRegistry[DELIVER_NAIVE] = deliverFactory.CreateNaiveRobot()
	RobotRegistry[DELIVER_INTERMEDIATE] = deliverFactory.CreateIntermediateRobot()
	RobotRegistry[DELIVER_ADVANCED] = deliverFactory.CreateAdvancedRobot()
}

func (reg robotRegistry) MakeCopy(category RobotCategory) robot.WorkableRobot {
	if reg[category] != nil {
		return reg[category].Clone()
	} else {
		return nil
	}
}
