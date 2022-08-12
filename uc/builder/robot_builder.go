package builder

import "design-pattern/entities/robot"

type RobotBuilder interface {
	BuildHands(r robot.WorkableRobot)
	BuildWheels(r robot.WorkableRobot)
	BuildCabins(r robot.WorkableRobot)
}
