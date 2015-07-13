#Edison Light Toys
##Background
Made for Gophercon 2015 Gobot Hackday (7/10/2015). Zero robotics experience beforehand.

##Compile Notes
For Edison, first make binary:

```GOARCH=386 GOOS=linux go build [go file]```

Then, send it to your Edison. 

```scp sphero root@[edison IP address]:/home/root/```

##Directory
* light_rainbow.go
* sphero.go

###Light Rainbow
This program needs a screen, a rotary sensor in port 0, a light sensor in port 1, and a temperature sensor in port 2. 
Note that the temperature in the room should be between 25 and 35 degrees Fahrenheit.
The rotary sensor maps to the red value, light to green, temperature to blue.

###Sphero
This needs an accelerometer and a sphero. Use the accelerometer (move in x and y axis) to move the sphero. 

###Temp Sensor
This is mainly for the Kill Room. Given a temp sensor (ON PORT 1!!) and a screen, displays the temperature in the room in F and C. The color of the screen changes from shades of blue (60) when it's hella cold to red (80) when it's hella hot.
