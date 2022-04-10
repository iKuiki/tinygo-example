package fonts_test

import (
	"fmt"
	"testing"
	"tinygo-example/stm32/4_1-oled/fonts"
)

// 此Test主要方便在电脑端查看点阵构型
func TestPrint(t *testing.T) {
	res := showChar('B')
	t.Log(res)
}

func showChar(Char byte) (res string) {
	// var x, y uint8 = 0, 0 //设置光标位置在上半部分
	fmt.Printf("\n1:\n")
	for _, d := range fonts.Font8x16[Char-' '][:8] {
		fmt.Printf("%x ", d)
		res += "|"
		for i := uint8(0); i < 8; i++ {
			if (d & (0x80 >> i)) > 0 {
				res += "*"
			} else {
				res += "_"
			}
		}
		res += "|\n"
	}

	fmt.Printf("\n2:\n")
	for _, d := range fonts.Font8x16[Char-' '][8:] {
		fmt.Printf("%x ", d)
		res += "|"
		for i := uint8(0); i < 8; i++ {
			if (d & (0x80 >> i)) > 0 {
				res += "*"
			} else {
				res += "_"
			}
		}
		res += "|\n"
	}
	fmt.Printf("\nend\n")
	return
}
