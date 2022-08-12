package hand

import "design-pattern/entities/robot/components"

type Hand struct {
	components.RobotComponent
	WorkableHand
}

type WorkableHand interface {
	RaiseUp()
	PutDown()
	Wave()
}
