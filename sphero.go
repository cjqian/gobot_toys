package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/i2c"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
	"github.com/hybridgroup/gobot/platforms/sphero"
	"math"
	"time"
)

func main() {
	gbot := gobot.NewGobot()

	adaptor := sphero.NewSpheroAdaptor("Sphero", "/dev/rfcomm0")
	board := edison.NewEdisonAdaptor("edison")
	spheroDriver := sphero.NewSpheroDriver(adaptor, "sphero")
	sensorAccel := i2c.NewGroveAccelerometerDriver(board, "accel")

	work := func() {
		gobot.Every(time.Millisecond*20, func() {
			if x, y, z, err := sensorAccel.XYZ(); err == nil {
				degree := math.Sin(y/x) * 360
				spheroDriver.Roll(uint8(100+z), uint16(degree))
			}
		})
	}

	robot := gobot.NewRobot("sphero",
		[]gobot.Connection{adaptor, board},
		[]gobot.Device{sensorAccel, spheroDriver},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
