package oled

import (
	"machine"
	"time"
)

const (
	// 屏幕i2c连接的针脚
	pinSCL = machine.PB8
	pinSDA = machine.PB9
	pinVCC = machine.PB7
	pinGND = machine.PB6
)

// 初始化oled的i2c
func i2cInit() {
	pinSCL.Configure(machine.PinConfig{Mode: machine.PinOutput50MHz})
	pinSDA.Configure(machine.PinConfig{Mode: machine.PinOutput50MHz})
	pinSCL.High()
	pinSDA.High()
}

// i2c开始
func i2cStart() {
	pinSDA.High()
	pinSCL.High()
	pinSDA.Low()
	pinSCL.Low()
}

// i2c停止
func i2cStop() {
	pinSDA.Low()
	pinSCL.High()
	pinSDA.High()
}

func i2cSendByte(b uint8) {
	for i := uint8(0); i < 8; i++ {
		pinSDA.Set((b & (0x80 >> i)) > 0)
		pinSCL.High()
		pinSCL.Low()
	}
	// 额外一个时钟信号，不处理应答
	pinSCL.High()
	pinSCL.Low()
}

func writeCommand(command uint8) {
	i2cStart()
	i2cSendByte(0x78) // 从机地址
	i2cSendByte(0x00) // 写命令
	i2cSendByte(command)
	i2cStop()
}

func writeData(data uint8) {
	i2cStart()
	i2cSendByte(0x78) // 从机地址
	i2cSendByte(0x40) // 写数据
	i2cSendByte(data)
	i2cStop()
}

func setCursor(y uint8, x uint8) {
	writeCommand(0xB0 | y)                 //设置Y位置
	writeCommand(0x10 | ((x & 0xF0) >> 4)) //设置X位置低4位
	writeCommand(0x00 | (x & 0x0F))        //设置X位置高4位
}

// Clear OLED清屏
// @brief  OLED清屏
// @param  无
// @retval 无
func Clear() {
	var i, j uint8
	for j = 0; j < 8; j++ {
		setCursor(j, 0)
		for i = 0; i < 128; i++ {
			writeData(0x00)
		}
	}
}

// ShowChar OLED显示一个字符
// @brief  OLED显示一个字符
// @param  Line 行位置，范围：1~4
// @param  Column 列位置，范围：1~16
// @param  Char 要显示的一个字符，范围：ASCII可见字符
// @retval 无
func ShowChar(Line uint8, Column uint8, Char byte) {
	var i uint8
	setCursor((Line-1)*2, (Column-1)*8) //设置光标位置在上半部分
	for i = 0; i < 8; i++ {
		writeData(font8x16[Char-' '][i]) //显示上半部分内容
	}
	setCursor((Line-1)*2+1, (Column-1)*8) //设置光标位置在下半部分
	for i = 0; i < 8; i++ {
		writeData(font8x16[Char-' '][i+8]) //显示下半部分内容
	}
}

// ShowString OLED显示字符串
// @brief  OLED显示字符串
// @param  Line 起始行位置，范围：1~4
// @param  Column 起始列位置，范围：1~16
// @param  String 要显示的字符串，范围：ASCII可见字符
// @retval 无
func ShowString(Line uint8, Column uint8, char string) {
	var i uint8
	for i = 0; i < uint8(len(char)); i++ {
		ShowChar(Line, Column+i, char[i])
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
func ShowNum(Line uint8, Column uint8, Number uint32, Length uint8) {
	var i uint8
	for i = 0; i < Length; i++ {
		// showChar(Line, Column+i, Number/pow(10, Length-i-1)%10+'0')
		ShowChar(Line, Column+i, byte(Number/pow(10, uint32(Length-i-1))%10+16))
	}
}

// ShowSignedNum OLED显示数字（十进制，带符号数）
// @brief  OLED显示数字（十进制，带符号数）
// @param  Line 起始行位置，范围：1~4
// @param  Column 起始列位置，范围：1~16
// @param  Number 要显示的数字，范围：-2147483648~2147483647
// @param  Length 要显示数字的长度，范围：1~10
// @retval 无
func ShowSignedNum(Line uint8, Column uint8, Number int32, Length uint8) {
	var i uint8
	var Number1 uint32
	if Number >= 0 {
		ShowChar(Line, Column, '+')
		Number1 = uint32(Number)
	} else {
		ShowChar(Line, Column, '-')
		Number1 = uint32(-Number)
	}
	for i = 0; i < Length; i++ {
		// showChar(Line, Column+i+1, Number1/pow(10, Length-i-1)%10+'0')
		ShowChar(Line, Column+i+1, byte(Number1/pow(10, uint32(Length-i-1))%10+16))
	}
}

// ShowHexNum OLED显示数字（十六进制，正数）
// @brief  OLED显示数字（十六进制，正数）
// @param  Line 起始行位置，范围：1~4
// @param  Column 起始列位置，范围：1~16
// @param  Number 要显示的数字，范围：0~0xFFFFFFFF
// @param  Length 要显示数字的长度，范围：1~8
// @retval 无
func ShowHexNum(Line uint8, Column uint8, Number uint32, Length uint8) {
	var i, SingleNumber uint8
	for i = 0; i < Length; i++ {
		SingleNumber = uint8(Number / pow(16, uint32(Length-i-1)) % 16)
		if SingleNumber < 10 {
			ShowChar(Line, Column+i, SingleNumber+'0')
		} else {
			ShowChar(Line, Column+i, SingleNumber-10+'A')
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
func ShowBinNum(Line uint8, Column uint8, Number uint32, Length uint8) {
	var i uint8
	for i = 0; i < Length; i++ {
		ShowChar(Line, Column+i, byte(Number/pow(2, uint32(Length-i-1))%2+'0'))
	}
}

// Init 初始化Oled屏幕
func Init() {
	// vcc上电，gnd接地
	pinVCC.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pinVCC.High()
	pinGND.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pinGND.Low()
	// 上电延时
	time.Sleep(time.Microsecond * 1000)

	// ============ 开始初始化 ============

	//端口初始化
	i2cInit()

	writeCommand(0xAE) //关闭显示

	writeCommand(0xD5) //设置显示时钟分频比/振荡器频率
	writeCommand(0x80)

	writeCommand(0xA8) //设置多路复用率
	writeCommand(0x3F)

	writeCommand(0xD3) //设置显示偏移
	writeCommand(0x00)

	writeCommand(0x40) //设置显示开始行

	writeCommand(0xA1) //设置左右方向，0xA1正常 0xA0左右反置

	writeCommand(0xC8) //设置上下方向，0xC8正常 0xC0上下反置

	writeCommand(0xDA) //设置COM引脚硬件配置
	writeCommand(0x12)

	writeCommand(0x81) //设置对比度控制
	writeCommand(0xCF)

	writeCommand(0xD9) //设置预充电周期
	writeCommand(0xF1)

	writeCommand(0xDB) //设置VCOMH取消选择级别
	writeCommand(0x30)

	writeCommand(0xA4) //设置整个显示打开/关闭

	writeCommand(0xA6) //设置正常/倒转显示

	writeCommand(0x8D) //设置充电泵
	writeCommand(0x14)

	writeCommand(0xAF) //开启显示

	Clear() //OLED清屏
}
