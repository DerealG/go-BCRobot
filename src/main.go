package main

import (
	"Nlog"
	"serial"
)

func main() {
	Nlog.Info.Println("Start")
	serial.Write("Hello from go\n")
	for {
		serial.Read()
	}
}
