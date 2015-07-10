/* Changes R, G, B as function of rotary sensor, light in room, and temperature. Prints to screen. */
/* Screen must connect to I2C, rot/light/temp to ports 0, 1, 2 respectively.*/
package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/i2c"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
	"time"
)

/*
This function is used if only the rotary port is in play.
func mapValueToColor(i int) (int, int, int) {
	// Masking to account for very sensitive rotary
	red := ((i & (1 << 1)) << 2) | ((i & (1 << 4)) << 0) | ((i & (1 << 7)) >> 2)
	green := ((i & (1 << 2)) << 1) | ((i & (1 << 5)) >> 1) | ((i & (1 << 8)) >> 3)
	blue := ((i & (1 << 3)) << 0) | ((i & (1 << 6)) >> 2) | ((i & (1 << 9)) >> 4)

	return red, green, blue
}
*/
func main() {
	gbot := gobot.NewGobot()

	board := edison.NewEdisonAdaptor("edison")
	screen := i2c.NewGroveLcdDriver(board, "screen")
	sensorRotary := gpio.NewGroveRotaryDriver(board, "sensor", "0")
	sensorLight := gpio.NewGroveLightSensorDriver(board, "light", "1")
	sensorTemp := gpio.NewGroveTemperatureSensorDriver(board, "temp", "2")

	work := func() {
		var r, g, b int

		gobot.On(sensorRotary.Event("data"), func(data interface{}) {
			r = data.(int) >> 2
		})

		gobot.On(sensorLight.Event("data"), func(data interface{}) {
			fmt.Printf("%d\n", data)
			g = data.(int) * 255 / 790
		})

		gobot.Every(time.Millisecond*500, func() {
			b = int(gobot.ToScale(gobot.FromScale(sensorTemp.Temperature(), 25, 35), 0, 255))
			screen.Clear()
			screen.Home()
			screen.SetRGB(r, g, b)
			screen.Write(fmt.Sprintf("#%02X%02X%02X", r, g, b))
		})

	}

	robot := gobot.NewRobot("screenBot",
		[]gobot.Connection{board},
		[]gobot.Device{screen, sensorRotary, sensorLight, sensorTemp},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
