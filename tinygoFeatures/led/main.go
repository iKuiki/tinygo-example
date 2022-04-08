package main

import (
	"machine"
	"time"
)

// 本程序驱动版载led做每秒2次闪烁

const (
	led = machine.LED
)

func init() { // 初始化LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func main() {
	// 此led为低电平点亮
	led.High()
	for { // 主程序下，循环切换led灯状态
		led.Low()
		time.Sleep(time.Millisecond * 100)

		led.High()
		time.Sleep(time.Millisecond * 100)

		led.Low()
		time.Sleep(time.Millisecond * 100)

		led.High()
		time.Sleep(time.Millisecond * 1400)
	}
}
