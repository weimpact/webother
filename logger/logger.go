package logger

import "log"

func Errorf(format string, args ...interface{}) {
	log.Println(format, args)
}

func Infof(format string, args ...interface{}) {
	log.Println(format, args)
}
