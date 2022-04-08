package main

import (
	"time"
	"tinygo-example/stm32/4_1-oled/oled"
)

func init() {
	oled.Init()
}

func main() {
	// 此实验无法通过编译，原因待查
	oled.ShowChar(1, 1, 'A')
	for {
		time.Sleep(time.Second)
	}
}
