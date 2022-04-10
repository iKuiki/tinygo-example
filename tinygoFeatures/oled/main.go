package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ssd1306"
)

const (
	led = machine.LED
)

var (
	screen ssd1306.Device
)

func init() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	screen = ssd1306.NewI2C(machine.I2C0)
	screen.Configure(ssd1306.Config{
		Address: 0x3C,
		Width:   128,
		Height:  64,
	})

	screen.ClearDisplay()
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

	var i uint8
	for {
		if i != 0 { // 通过设置屏幕的buffer来改变屏幕显示内容
			var b []byte // 设置为屏幕内容的byte切片，此byte切片的长度必须为1024
			for i := 0; i < 1024; i++ {
				b = append(b, byte(i%255))
			}
			screen.SetBuffer(b)
			screen.Display()

			i = 0
		} else { // 通过设置像素点来改变屏幕内容
			screen.ClearBuffer() // 先清除屏幕
			// 尝试通过像素来控制屏幕显示字母K
			screen.SetPixel(43, 32, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(43, 33, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(43, 34, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(43, 35, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(43, 36, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(43, 37, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(43, 38, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(43, 39, color.RGBA{255, 255, 255, 255})

			screen.SetPixel(46, 32, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(46, 33, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(45, 33, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(45, 34, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(44, 34, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(44, 35, color.RGBA{255, 255, 255, 255})

			screen.SetPixel(44, 36, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(45, 37, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(46, 37, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(47, 37, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(47, 38, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(48, 38, color.RGBA{255, 255, 255, 255})
			screen.SetPixel(48, 39, color.RGBA{255, 255, 255, 255})

			screen.Display()

			i = 1
		}

		time.Sleep(time.Second)

	}

}
