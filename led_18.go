package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
 
 
   fmt.Println("Hello \n")
   
   // 모든 pin은 bcm2835 기준입니다.
   
 

	// Get the pin for each of the lights
	redPin := rpio.Pin(18)
 

	// Set the pins to output mode
	redPin.Output()
	 

	// Clean up on ctrl-c and turn lights out
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		redPin.Low() 
   os.Exit(0)
	 
	}()

	defer rpio.Close()

	// Turn lights off to start.
	redPin.Low()
	 
	// A while true loop.
	for {
		// Red
		redPin.High()
		time.Sleep(time.Second * 1)
	  redPin.Low() 
 		time.Sleep(time.Second * 1)

	 
	}
}
