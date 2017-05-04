package tool

import (
	"fmt"

	"github.com/tokitgit/tokit/tokit_go_src/go-colortext"
)

type Errvec_t []string

func Echo_errvec(errs *Errvec_t) {
	for _, v := range *errs {
		EchoErr("  %s", v)
	}
}

func GetErrMsg(cmd string, arg ...interface{}) string {
	return fmt.Sprintf(cmd, arg...)
}

func EchoErr(cmd string, arg ...interface{}) {
	gocolortext.ChangeColor(gocolortext.Red, true, gocolortext.None, false)
	fmt.Println(fmt.Sprintf(cmd, arg...))
}

func EchoWarn(cmd string, arg ...interface{}) {
	gocolortext.ChangeColor(gocolortext.Yellow, true, gocolortext.None, false)
	fmt.Println(fmt.Sprintf(cmd, arg...))
}

func EchoOk(cmd string, arg ...interface{}) {
	gocolortext.ChangeColor(gocolortext.Green, true, gocolortext.None, false)
	fmt.Println(fmt.Sprintf(cmd, arg...))
}

func ResetConsoleColor() {
	gocolortext.ResetColor()
}
