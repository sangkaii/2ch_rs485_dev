package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"os/signal"
	"syscall"
	"time"
 
  

	"github.com/goburrow/modbus"
 
)


func check_file(e error) {
	if e != nil {
		//panic(e)
	}
}

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
 
 
   fmt.Println("Hello \n")
   
   // 모든 pin은 bcm2835 기준입니다.
   
 

	// Get the pin for each of the lights
	redPin := rpio.Pin(26)
 


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
   	func_relay("on")
		time.Sleep(time.Second * 1)
	  redPin.Low() 
   func_relay("off")
 		time.Sleep(time.Second * 1)

	 
	}
}



/*

	// function name
	// Bit access
	ReadDiscreteInputs = 2
	ReadCoils          = 1
	WriteSingleCoil    = 5
	WriteMultipleCoils = 15

	// 16-bit access
	ReadInputRegisters         = 4
	ReadHoldingRegisters       = 3
	WriteSingleRegister        = 6
	WriteMultipleRegisters     = 16
	ReadWriteMultipleRegisters = 23
	MaskWriteRegister          = 22
	ReadFIFOQueue              = 24

*/

/*

    // Slave ID : 1 일때 ------
	Relay 0-ON : 01 05 00 00 FF 00
	Relay 0-OFF: 01 05 00 00 00 00

	Relay 1-ON : 01 05 00 01 FF 00
	Relay 1-OFF: 01 05 00 01 00 00


*/

func func_relay(state string) {
	// Modbus RTU/ASCII
	//handler := modbus.NewRTUClientHandler("/dev/ttyUSB0")
	handler := modbus.NewRTUClientHandler("/dev/ttySC1")
 
 
 
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.Timeout = 1 * time.Second

	handler.SlaveId = 1

	var results []byte
	var err error

	err = handler.Connect()
	defer handler.Close()

/*
Pin connection:
Pin connections can be viewed in lib Config DEV_Config.C. Here again:
2-CH_RS485_HAT      =>    RPI(BCM)
VCC                 ->    3.3V/5V
GND                 ->    GND
SCLK                ->    21
MISO                ->    20
MOSI                ->    19
CS                  ->    18
IRQ                 ->    24
EN1                 ->    27
EN2                 ->    22

*/

 txden_1_pin := rpio.Pin(27)
 txden_1_pin.Output()
 time.Sleep(time.Second * 2)

  
 
  txden_1_pin.Low()
  time.Sleep(time.Millisecond * 5)
  
	if err != nil {
		fmt.Println("Serial Port not Started !!!")
		return
	}

	client := modbus.NewClient(handler)

	fmt.Println("Action : ", state)

	if state == "off" {
		results, err = client.WriteSingleCoil(0, 0x0000)
	}

	if state == "on" {
 		results, err = client.WriteSingleCoil(0, 0xff00)
	}

	if err != nil {
		fmt.Println(err)
	}
 
 

 time.Sleep(time.Second * 2)
 
 txden_1_pin.High()
 
 time.Sleep(time.Second * 2)
 time.Sleep(time.Millisecond * 5)
  
	UNUSED (results, err)

}

func UNUSED(x ...interface{}) {}
