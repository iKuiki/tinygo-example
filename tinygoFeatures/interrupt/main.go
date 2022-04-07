package main

import (
	"machine"
	"time"
)

// 此demo使用pa0针脚的电平改变的中断，来切换板载led灯的状态，来演示中断的使用

const ( // 声明用到的pin
	// 板载led灯
	led = machine.LED
	// pa0用做button
	button = machine.PA0
)

func main() {
	// tinygo对外设初始化做了简化，在设置时，就会开启时钟

	// led设置为out，正常来说在C中这里应当设置为开漏输出，但是tinygo在这里不能设置为开漏（也不能设置为推挽），否则会遇到问题
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// 按键设置为拉低input
	button.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	button.SetInterrupt(machine.PinToggle, func(p machine.Pin) {
		if led.Get() {
			led.Low()
			// 这里绝对不能用sleep，否则直接卡死
			// 可能是sleep时，中断没有清除？
			// time.Sleep(time.Millisecond * 200)
		} else {
			led.High()
			// 这里绝对不能用sleep，否则直接卡死
			// 可能是sleep时，中断没有清除？
			// time.Sleep(time.Millisecond * 200)
		}
	})
	// 此led为下
	led.High()
	for { // 主程序下，循环切换led灯状态
		if led.Get() {
			led.Low()
		} else {
			led.High()
		}
		time.Sleep(time.Millisecond * 500)
	}
}
