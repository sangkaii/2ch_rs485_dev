package main

import (
    "fmt"
    "time"
    rpio "github.com/stianeikeland/go-rpio"
)

const (
    PWM_PIN = 18
)

func main() {
    err := rpio.Open()
    if err != nil {
        panic(err)
    }

    motor_pin_pwm := rpio.Pin(PWM_PIN)
    motor_pin_pwm.Mode(rpio.Pwm)
    motor_pin_pwm.Freq(1000)
    motor_pin_pwm.DutyCycle(0, 100)

    fmt.Println("Waiting...")
    time.Sleep(5 * time.Second)
    fmt.Println("Start the increase...")

    for i := 1; i <= 100; i++ {
        fmt.Println(i)
        motor_pin_pwm.DutyCycle(uint32(i), 100)
        time.Sleep(100 * time.Millisecond)
    }
    time.Sleep(5 * time.Second)
}
