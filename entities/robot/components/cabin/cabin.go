package cabin

import "design-pattern/entities/robot/components"

type Cabin struct {
	components.RobotComponent
	WorkableCabin
}

type WorkableCabin interface {
	Open()
	Close()
	Load()
	Unload()
}
