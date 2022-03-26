package utils

import "os"

func Catch(err error, cb func(errMessage string)) {
	if err != nil {
		cb(err.Error())
	}
}

// InterceptErrorsAndKillProcessImmediately 拦截错误且马上杀掉进程
func InterceptErrorsAndKillProcessImmediately(err error, cb func(msg string)) {
	if err != nil {
		BeforeStoppingProcess(func() {
			cb(err.Error())
		})
	}
}

// ColdKiller 冷漠的杀手 拦截错误不提供处理函数
func ColdKiller(err error) {
	if err != nil {
		BeforeStoppingProcess(func() {
			RedTips(
				err.Error(),
			)
		})
	}
}
func BeforeStoppingProcess(cb func()) {
	cb()
	os.Exit(1)
}
