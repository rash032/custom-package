package mylogger

import "log"

func LogInfo(message string){
	log.Println("INFO - %v", message)
}