// Open the seed studio grove kit box and find the little green bag labeled "Grove - Temperature sensor", open the
// bag and plug the grove connector cable into the grove slot labeled "A0".
package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/i2c"
	"github.com/hybridgroup/gobot/platforms/intel-iot/edison"
	"math"
	"time"
)

//assume minimum is 60, max is 80
func getRed(ftemp float64) int {
	//range from 0 to 20 now
	ftempConv := ftemp - 60
	red := (ftempConv / 20) * 255
	return int(red)
}

func getBlue(ftemp float64) int {
	ftempConv := ftemp - 60
	blue := (1 - (ftempConv / 20)) * 255
	return int(blue)
}

func main() {
	gbot := gobot.NewGobot()

	board := edison.NewEdisonAdaptor("board")
	screen := i2c.NewGroveLcdDriver(board, "screen")
	sensor := gpio.NewGroveTemperatureSensorDriver(board, "sensor", "1")

	work := func() {
		gobot.Every(time.Second*2, func() {
			screen.Clear()
			screen.Home()
			celsius := sensor.Temperature()
			fahrenheit := (celsius * 1.8) + 32
			screen.Write(fmt.Sprintf("%.2fF, %.2fC", fahrenheit, celsius))
			//set to params
			newFahrenheit := math.Max(60, math.Min(80, fahrenheit))
			screen.SetRGB(getRed(newFahrenheit), 0, getBlue(newFahrenheit))
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{board},
		[]gobot.Device{screen, sensor},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
