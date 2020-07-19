package fprint

import (
	"fmt"
)

var (
	greenBg   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	whiteBg   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellowBg  = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	redBg     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blueBg    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magentaBg = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyanBg    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	green     = string([]byte{27, 91, 51, 50, 109})
	white     = string([]byte{27, 91, 51, 55, 109})
	yellow    = string([]byte{27, 91, 51, 51, 109})
	red       = string([]byte{27, 91, 51, 49, 109})
	blue      = string([]byte{27, 91, 51, 52, 109})
	magenta   = string([]byte{27, 91, 51, 53, 109})
	cyan      = string([]byte{27, 91, 51, 54, 109})
	reset     = string([]byte{27, 91, 48, 109})
)

//  =============================    bg    ==========================================

func GreenBg(str ...interface{}) {
	fmt.Println(greenBg, str, reset)
}

func GrayBg(str ...interface{}) {
	fmt.Println(whiteBg, str, reset)
}

func YellowBg(str ...interface{}) {
	fmt.Println(yellowBg, str, reset)
}

func RedBg(str ...interface{}) {
	fmt.Println(redBg, str, reset)
}

func BlueBg(str ...interface{}) {
	fmt.Println(blueBg, str, reset)
}

func MagentaBg(str ...interface{}) {
	fmt.Println(magentaBg, str, reset)
}

func CyanBg(str ...interface{}) {
	fmt.Println(cyanBg, str, reset)
}

//  =============================    text    ==========================================

func Green(str ...interface{}) {
	fmt.Println(green, str, reset)
}
func White(str ...interface{}) {
	fmt.Println(white, str, reset)
}
func Yellow(str ...interface{}) {
	fmt.Println(yellow, str, reset)
}
func Red(str ...interface{}) {
	fmt.Println(red, str, reset)
}
func Blue(str ...interface{}) {
	fmt.Println(blue, str, reset)
}
func Magenta(str ...interface{}) {
	fmt.Println(magenta, str, reset)
}
func Cyan(str ...interface{}) {
	fmt.Println(cyan, str, reset)
}
