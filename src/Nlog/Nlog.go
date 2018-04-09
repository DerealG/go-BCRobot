package Nlog

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Net   *log.Logger
	Error *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "Info:", log.LstdFlags)
	Net = log.New(os.Stdout, "Net:", log.LstdFlags)
	Error = log.New(os.Stdout, "Error:", log.LstdFlags)
}
