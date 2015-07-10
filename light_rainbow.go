/* Changes R, G, B as function of rotary sensor, light in room, and temperature. Prints to screen. */
// Open the seed studio grove kit box and find the little green bag labeled "Grove - LCD", open the
// bag and plug the grove connector cable into the grove slot labeled "I2C".
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
func mapValueToColor(i int) (int, int, int) {
	//mask the 10bit input to an 8bit input
	//eightBit := i & 1111111110
	//fmt.Println("%d\n", eightBit)
	//green
	//blue
	red := ((i & (1 << 1)) << 2) | ((i & (1 << 4)) << 0) | ((i & (1 << 7)) >> 2)
	green := ((i & (1 << 2)) << 1) | ((i & (1 << 5)) >> 1) | ((i & (1 << 8)) >> 3)
	blue := ((i & (1 << 3)) << 0) | ((i & (1 << 6)) >> 2) | ((i & (1 << 9)) >> 4)
	//red, green, blue := 0, 0, 0
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
		//for light

		gobot.On(sensorRotary.Event("data"), func(data interface{}) {
			r = data.(int) >> 2
			//	r, g, b = mapValueToColor(data.(int))
		})

		gobot.On(sensorLight.Event("data"), func(data interface{}) {
			fmt.Printf("%d\n", data)
			g = data.(int) * 255 / 790
			//  r, g, b = mapValueToColor(data.(int))
		})

		gobot.Every(time.Millisecond*500, func() {
			b = int(gobot.ToScale(gobot.FromScale(sensorTemp.Temperature(), 25, 35), 0, 255))
			//	fmt.Printf("Temperature:%d\n", sensorTemp.Temperature())
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
