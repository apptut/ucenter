package logger

import (
	"log"
)

func Fatal(msg error) {
	log.Fatalf("[fatal]\t%s\n", msg.Error())
}

func Error(msg error) {
	log.Printf("[error]\t%s\n", msg.Error())
}

func Info(msg string) {
	log.Printf("[info]\t%s\n", msg)
}
