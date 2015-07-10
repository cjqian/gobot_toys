// Open the seed studio grove kit box and find the little green bag labeled "Grove - Piezo Vibration  Sensor", open the
// bag and plug the grove connector cable into the grove slot labeled "A0". 
package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
)

func main() {
	gbot := gobot.NewGobot()

	board := edison.NewEdisonAdaptor("edison")
	sensor := gpio.NewGrovePiezoVibrationSensorDriver(board, "sensor", "0")

	work := func() {
		gobot.On(sensor.Event(gpio.Vibration), func(data interface{}) {
			fmt.Println("got one!")
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{board},
		[]gobot.Device{sensor},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
