// Open the seed studio grove kit box and find the little green bag labeled "Grove - Touch Sensor", open the
// bag and plug the grove connector cable into the grove slot labeled "D2".
package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
)

func main() {
	gbot := gobot.NewGobot()

	e := edison.NewEdisonAdaptor("edison")
	touch := gpio.NewGroveTouchDriver(e, "touch", "2")

	work := func() {
		gobot.On(touch.Event(gpio.Push), func(data interface{}) {
			fmt.Println("On!")
		})

		gobot.On(touch.Event(gpio.Release), func(data interface{}) {
			fmt.Println("Off!")
		})

	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{e},
		[]gobot.Device{touch},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
