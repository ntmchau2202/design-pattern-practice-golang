package wheel

import "design-pattern/entities/robot/components"

type Wheel struct {
	components.RobotComponent
	WorkableWheel
}

type WorkableWheel interface {
	Run()
}
