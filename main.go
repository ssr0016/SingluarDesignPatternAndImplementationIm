package main

import (
	"fmt"
	"time"

)

// singular pattern

type logger struct {
	level    int8
	initFlag bool
}

var logMaster = &logger{}

func InitialiseLogger(level int8) *logger {
	if !logMaster.initFlag {
		logMaster = &logger{
			level:    level,
			initFlag: true,
		}
	}
	return logMaster
}

func (l *logger) Error(msg string) {
	if l.level >= 3 {
		fmt.Println("level: ", l.level)
		fmt.Println(msg, time.Now())
	}
}

func main() {
	log1 := InitialiseLogger(3)
	log2 := InitialiseLogger(4)
	log3 := InitialiseLogger(5)
	log4 := InitialiseLogger(6)

	log1.Error("message-1")
	log2.Error("message-2")
	log3.Error("message-3")
	log4.Error("message-4")

}
