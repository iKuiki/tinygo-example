package main

import (
	"machine"
	"time"
)

// 本程序驱动pa0～pa3上接的4个led灯做循环亮灯
// 原实验是8个led灯，此处简化为4个
// 原实验用的是对GPIOA寄存器直接写值，此处因无法直接访问寄存器，所以使用led对象

const (
	led1 = machine.PA0
	led2 = machine.PA1
	led3 = machine.PA2
	led4 = machine.PA3
)

func init() { // 初始化LED
	// 在tinygo下貌似当前只有output有效
	led1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led1.High() // 初始化关灯
	led2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led2.High() // 初始化关灯
	led3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led3.High() // 初始化关灯
	led4.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led4.High() // 初始化关灯
}

func main() {
	// 此led为低电平点亮
	for { // 主程序下，循环切换led灯状态
		led4.High()
		led1.Low()
		time.Sleep(time.Millisecond * 500)

		led1.High()
		led2.Low()
		time.Sleep(time.Millisecond * 500)

		led2.High()
		led3.Low()
		time.Sleep(time.Millisecond * 500)

		led3.High()
		led4.Low()
		time.Sleep(time.Millisecond * 500)
	}
}
