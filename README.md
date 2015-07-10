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
* light_rainbow


###Light Rainbow
This program needs a screen, a rotary sensor in port 0, a light sensor in port 1, and a temperature sensor in port 2. 
Note that the temperature in the room should be between 25 and 35 degrees Fahrenheit.
The rotary sensor maps to the red value, light to green, temperature to blue.
