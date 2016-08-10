//日志处理封装
package logger

import (
	"log"
)

//调试信息
func Debug(msg ...error) {
	for _, m := range msg {
		log.Print(m)
	}
}

//系统运行过程中错误，不影响系统的继续运行
func Info() {

}
