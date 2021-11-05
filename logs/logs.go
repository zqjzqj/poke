package logs

import (
	"github.com/gookit/color"
	"github/zqjzqj/poker/global"
	"log"
	"os"
	"time"
)

var IsPrintLog = true
var IsDebugPrint = true

//这个包用来统一的日志输出处理
//目前只做简单两个方法 后续根据具体需要在这里增加日志操作
func Println(v ...interface{}) {
	if IsPrintLog {
		log.Println(v...)
	}
}

func print2(color2 color.Color, v ...interface{}) {
	s := time.Now().Format(global.DateTimeFormatStr)
	v = append([]interface{}{"[" + s + "]"}, v...)
	color2.Light().Println(v...)
}

func PrintDebug(v ...interface{}) {
	if IsDebugPrint {
		print2(color.Cyan, v...)
	}
}

func PrintDebugErr(v ...interface{}) {
	if IsDebugPrint {
		print2(color.FgLightRed, v...)
	}
}
func PrintlnSuccess(v ...interface{}) {
	if IsPrintLog {
		print2(color.Green, v...)
	}
}

func PrintlnInfo(v ...interface{}) {
	if IsPrintLog {
		print2(color.LightCyan, v...)
	}
}

func PrintlnWarning(v ...interface{}) {
	if IsPrintLog {
		print2(color.Yellow, v...)
	}
}

func PrintErr(v ...interface{}) {
	if IsPrintLog {
		print2(color.FgLightRed, v...)
	}
}

func Fatal(v ...interface{}) {
	PrintErr(v...)
	os.Exit(1)
}
