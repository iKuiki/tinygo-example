package main

import (
	"machine"
	"time"
	oled "tinygo-example/stm32/4_1-oled/oled2"
)

var screen oled.Screen

const (
	led = machine.LED
)

func init() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	screen = oled.NewScreen()
}

func main() {
	// 设置led状态灯
	led.High()  // 此led为低电平点亮，将其设置为高电平以熄灭
	go func() { // 后台运行led状态灯
		for {
			led.Low()
			time.Sleep(time.Millisecond * 100)

			led.High()
			time.Sleep(time.Millisecond * 100)

			led.Low()
			time.Sleep(time.Millisecond * 100)

			led.High()
			time.Sleep(time.Millisecond * 1400)
		}
	}()

	screen.ShowChar(1, 1, 'A')
	screen.ShowString(1, 3, "HelloWorld!")
	// 以下注释的方法，还需要再理解理解，会导致死机
	// screen.ShowNum(2, 1, 12345, 5)
	// screen.ShowSignedNum(2, 7, -66, 2)
	// screen.ShowHexNum(3, 1, 0xAA55, 4)
	// screen.ShowBinNum(4, 1, 0xAA55, 16)
	for {
		time.Sleep(time.Second)
	}
}
