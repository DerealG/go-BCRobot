package serial

import (
	"Nlog"
	"github.com/huin/goserial"
	"io"
)

var gReadWriteCloser io.ReadWriteCloser

func init() {
	Nlog.Info.Println("serial init.")
	c := &goserial.Config{Name: "/dev/ttyUSB0", Baud: 115200}
	closer, err := goserial.OpenPort(c)
	if err != nil {
		Nlog.Error.Println(err)
	} else {
		gReadWriteCloser = closer
	}
}

func Write(writeBuff string) int {
	n, err := gReadWriteCloser.Write([]byte(writeBuff))
	if err != nil {
		Nlog.Error.Println("serial write error :", err)
		return -1
	} else {
		Nlog.Net.Println("[W] ", writeBuff)
		return n
	}
}

func Read() string {
	buf := make([]byte, 1024)
	n, err := gReadWriteCloser.Read(buf)
	if err != nil {
		if err != io.EOF { //EOF error isn't a real error.
			Nlog.Error.Println("serial read error :", err)
		}
		return ""
	} else {
		readBuff := string(buf[:n])
		Nlog.Net.Println("[R] ", readBuff)
		return readBuff
	}
}
