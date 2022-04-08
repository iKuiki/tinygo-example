package main

import (
	"machine"
	"time"
)

// 本程序驱动PA0上的led做每秒1次闪烁

const (
	led = machine.PA0
)

func init() { // 初始化LED
	// 在tinygo下貌似当前只有output有效
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func main() {
	// 此led为低电平点亮
	led.High()
	for { // 主程序下，循环切换led灯状态
		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)

		led.Set(false)
		time.Sleep(time.Millisecond * 500)

		led.Set(true)
		time.Sleep(time.Millisecond * 500)
	}
}
