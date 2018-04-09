package main

import (
	"Nlog"
	"serial"
	"time"
)

func main() {
	Nlog.Info.Println("Start")
	go ReadThd()
	for {
		serial.Write("I'm here.\r\n")
		time.Sleep(5 * time.Second)
	}
}

func ReadThd() {
	for {
		serial.Read()
	}
}
