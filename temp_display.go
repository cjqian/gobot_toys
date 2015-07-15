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

var (
	min  float64
	max  float64
	diff float64
)

//assume minimum is 60, max is 80
func getRed(ftemp float64) int {
	//range from 0 to 20 now
	ftempConv := ftemp - min
	red := (ftempConv / diff) * 255
	return int(red)
}

func getBlue(ftemp float64) int {
	ftempConv := ftemp - min
	blue := (1 - (ftempConv / diff)) * 255
	return int(blue)
}

func main() {
	gbot := gobot.NewGobot()

	board := edison.NewEdisonAdaptor("board")
	screen := i2c.NewGroveLcdDriver(board, "screen")
	buzzer := gpio.NewBuzzerDriver(board, "buzzer", "3")
	sensor := gpio.NewGroveTemperatureSensorDriver(board, "sensor", "1")

	//the hot song is
	work := func() {
		//initial calibration in fahrenheit
		min = 60.0
		max = 80.0
		diff = max - min
		type note struct {
			tone     float64
			duration float64
		}
		//the cold song is "do you want to build a snowman"
		coldSong := []note{
			{gpio.C3, gpio.Quarter},
			{gpio.C3, gpio.Quarter},
			{gpio.C3, gpio.Quarter},
			{gpio.G3, gpio.Quarter},
			{gpio.C3, gpio.Quarter},
			{gpio.E4, gpio.Quarter},
			{gpio.D4, gpio.Half},
			{gpio.E4, gpio.Half},
		}
		//the hot song is "let's get it started in here"
		hotSong := []note{
			{gpio.E4, gpio.Quarter},
			{gpio.E4, gpio.Quarter},
			{gpio.E4, gpio.Quarter},
			{gpio.E4, gpio.Quarter},
			{gpio.C3, gpio.Quarter},
			{gpio.A3, gpio.Quarter},
			{gpio.C3, gpio.Quarter},
		}
		gobot.Every(time.Second, func() {
			screen.Clear()
			time.Sleep(5 * time.Millisecond)
			screen.Home()
			time.Sleep(5 * time.Millisecond)

			//get temps
			celsius := sensor.Temperature()
			fahrenheit := (celsius * 1.8) + 32
			//recalibrate
			min = math.Min(fahrenheit, min)
			max = math.Max(fahrenheit, max)
			diff = max - min

			//check
			if fahrenheit > 80 {
				for _, val := range hotSong {
					buzzer.Tone(val.tone, val.duration)
					<-time.After(10 * time.Millisecond)
				}
				screen.Write(fmt.Sprintf("TOO HOT (%.2gF)", fahrenheit))
			} else if fahrenheit < 60 {
				for _, val := range coldSong {
					buzzer.Tone(val.tone, val.duration)
					<-time.After(10 * time.Millisecond)
				}

				screen.Write(fmt.Sprintf("TOO COLD (%.2gF)", fahrenheit))
			} else {
				screen.Write(fmt.Sprintf("%.2gF, %.2gC", fahrenheit, celsius))
			}
			time.Sleep(5 * time.Millisecond)
			fmt.Printf("%.2gF, %.2gC\n", fahrenheit, celsius)

			//set to params
			screen.SetRGB(getRed(fahrenheit), 0, getBlue(fahrenheit))
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{board},
		[]gobot.Device{screen, buzzer, sensor},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
