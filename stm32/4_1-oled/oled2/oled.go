package oled

import (
	"image/color"
	"machine"
	"tinygo-example/stm32/4_1-oled/fonts"

	"tinygo.org/x/drivers/ssd1306"
)

// 因为直接使用gpio针脚模拟i2c通信的写法会导致编译后文件超过空间
// 所以此处暂时使用板载i2c模块来完成i2c通信

// Screen TTF屏幕
type Screen interface {
	// 清空屏幕内容
	ClearDisplay()

	// ShowChar OLED显示一个字符
	// @brief  OLED显示一个字符
	// @param  Line 行位置，范围：1~4
	// @param  Column 列位置，范围：1~16
	// @param  Char 要显示的一个字符，范围：ASCII可见字符
	// @retval 无
	ShowChar(Line uint8, Column uint8, Char byte)

	// ShowString OLED显示字符串
	// @brief  OLED显示字符串
	// @param  Line 起始行位置，范围：1~4
	// @param  Column 起始列位置，范围：1~16
	// @param  String 要显示的字符串，范围：ASCII可见字符
	// @retval 无
	ShowString(Line uint8, Column uint8, char string)

	// ShowNum OLED显示数字（十进制，正数）
	// @brief  OLED显示数字（十进制，正数）
	// @param  Line 起始行位置，范围：1~4
	// @param  Column 起始列位置，范围：1~16
	// @param  Number 要显示的数字，范围：0~4294967295
	// @param  Length 要显示数字的长度，范围：1~10
	// @retval 无
	ShowNum(Line uint8, Column uint8, Number uint32, Length uint8)

	// ShowSignedNum OLED显示数字（十进制，带符号数）
	// @brief  OLED显示数字（十进制，带符号数）
	// @param  Line 起始行位置，范围：1~4
	// @param  Column 起始列位置，范围：1~16
	// @param  Number 要显示的数字，范围：-2147483648~2147483647
	// @param  Length 要显示数字的长度，范围：1~10
	// @retval 无
	ShowSignedNum(Line uint8, Column uint8, Number int32, Length uint8)

	// ShowHexNum OLED显示数字（十六进制，正数）
	// @brief  OLED显示数字（十六进制，正数）
	// @param  Line 起始行位置，范围：1~4
	// @param  Column 起始列位置，范围：1~16
	// @param  Number 要显示的数字，范围：0~0xFFFFFFFF
	// @param  Length 要显示数字的长度，范围：1~8
	// @retval 无
	ShowHexNum(Line uint8, Column uint8, Number uint32, Length uint8)

	// ShowBinNum OLED显示数字（二进制，正数）
	// @brief  OLED显示数字（二进制，正数）
	// @param  Line 起始行位置，范围：1~4
	// @param  Column 起始列位置，范围：1~16
	// @param  Number 要显示的数字，范围：0~1111 1111 1111 1111
	// @param  Length 要显示数字的长度，范围：1~16
	// @retval 无
	ShowBinNum(Line uint8, Column uint8, Number uint32, Length uint8)
}

type screen struct {
	display ssd1306.Device
}

// NewScreen 初始化屏幕参数
func NewScreen() Screen {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
		// SCL:       machine.PB6, // 使用板载i2c0针脚
		// SDA:       machine.PB7, // 使用板载i2c0针脚
	})

	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: 0x3C,
		Width:   128,
		Height:  64,
	})

	// 初始化清屏
	display.ClearDisplay()

	return &screen{
		display: display,
	}
}

// 清空屏幕内容
func (scr *screen) ClearDisplay() {
	scr.display.ClearDisplay()
}

// ShowChar OLED显示一个字符
// @brief  OLED显示一个字符
// @param  Line 行位置，范围：1~4
// @param  Column 列位置，范围：1~16
// @param  Char 要显示的一个字符，范围：ASCII可见字符
// @retval 无
func (scr *screen) ShowChar(Line uint8, Column uint8, Char byte) {
	x, y := (Column-1)*8, (Line)*16-8 //设置光标位置在上半部分
	for _, d := range fonts.Font8x16[Char-' '][:8] {
		for i := uint8(0); i < 8; i++ {
			var r color.RGBA
			if (d & (0x80 >> i)) > 0 {
				r = color.RGBA{255, 255, 255, 255}
			}
			scr.display.SetPixel(int16(x), int16(y-i), r)
		}
		x++
	}

	x, y = (Column-1)*8, (Line)*16 //设置光标位置在下半部分

	for _, d := range fonts.Font8x16[Char-' '][8:] {
		for i := uint8(0); i < 8; i++ {
			var r color.RGBA
			if (d & (0x80 >> i)) > 0 { // 用该字段比上debug
				r = color.RGBA{255, 255, 255, 255}
			}
			scr.display.SetPixel(int16(x), int16(y-i), r)
		}
		x++
	}
	scr.display.Display()
}

// ShowString OLED显示字符串
// @brief  OLED显示字符串
// @param  Line 起始行位置，范围：1~4
// @param  Column 起始列位置，范围：1~16
// @param  String 要显示的字符串，范围：ASCII可见字符
// @retval 无
func (scr *screen) ShowString(Line uint8, Column uint8, char string) {
	var i uint8
	for i = 0; i < uint8(len(char)); i++ {
		scr.ShowChar(Line, Column+i, char[i])
	}
}

/**
 * @brief  OLED次方函数
 * @retval 返回值等于X的Y次方
 */
func pow(X uint32, Y uint32) uint32 {
	Result := uint32(1)
	for Y > 0 {
		Y--
		Result *= X
	}
	return Result
}

// ShowNum OLED显示数字（十进制，正数）
// @brief  OLED显示数字（十进制，正数）
// @param  Line 起始行位置，范围：1~4
// @param  Column 起始列位置，范围：1~16
// @param  Number 要显示的数字，范围：0~4294967295
// @param  Length 要显示数字的长度，范围：1~10
// @retval 无
func (scr *screen) ShowNum(Line uint8, Column uint8, Number uint32, Length uint8) {
	var i uint8
	for i = 0; i < Length; i++ {
		// showChar(Line, Column+i, Number/pow(10, Length-i-1)%10+'0')
		scr.ShowChar(Line, Column+i, byte(Number/pow(10, uint32(Length-i-1))%10+16))
	}
}

// ShowSignedNum OLED显示数字（十进制，带符号数）
// @brief  OLED显示数字（十进制，带符号数）
// @param  Line 起始行位置，范围：1~4
// @param  Column 起始列位置，范围：1~16
// @param  Number 要显示的数字，范围：-2147483648~2147483647
// @param  Length 要显示数字的长度，范围：1~10
// @retval 无
func (scr *screen) ShowSignedNum(Line uint8, Column uint8, Number int32, Length uint8) {
	var i uint8
	var Number1 uint32
	if Number >= 0 {
		scr.ShowChar(Line, Column, '+')
		Number1 = uint32(Number)
	} else {
		scr.ShowChar(Line, Column, '-')
		Number1 = uint32(-Number)
	}
	for i = 0; i < Length; i++ {
		// showChar(Line, Column+i+1, Number1/pow(10, Length-i-1)%10+'0')
		scr.ShowChar(Line, Column+i+1, byte(Number1/pow(10, uint32(Length-i-1))%10+16))
	}
}

// ShowHexNum OLED显示数字（十六进制，正数）
// @brief  OLED显示数字（十六进制，正数）
// @param  Line 起始行位置，范围：1~4
// @param  Column 起始列位置，范围：1~16
// @param  Number 要显示的数字，范围：0~0xFFFFFFFF
// @param  Length 要显示数字的长度，范围：1~8
// @retval 无
func (scr *screen) ShowHexNum(Line uint8, Column uint8, Number uint32, Length uint8) {
	var i, SingleNumber uint8
	for i = 0; i < Length; i++ {
		SingleNumber = uint8(Number / pow(16, uint32(Length-i-1)) % 16)
		if SingleNumber < 10 {
			scr.ShowChar(Line, Column+i, SingleNumber+'0')
		} else {
			scr.ShowChar(Line, Column+i, SingleNumber-10+'A')
		}
	}
}

// ShowBinNum OLED显示数字（二进制，正数）
// @brief  OLED显示数字（二进制，正数）
// @param  Line 起始行位置，范围：1~4
// @param  Column 起始列位置，范围：1~16
// @param  Number 要显示的数字，范围：0~1111 1111 1111 1111
// @param  Length 要显示数字的长度，范围：1~16
// @retval 无
func (scr *screen) ShowBinNum(Line uint8, Column uint8, Number uint32, Length uint8) {
	var i uint8
	for i = 0; i < Length; i++ {
		scr.ShowChar(Line, Column+i, byte(Number/pow(2, uint32(Length-i-1))%2+'0'))
	}
}
