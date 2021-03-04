package pelcod

import (
	"fmt"
	"io"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

var myport io.ReadWriteCloser

func OpenSerialPort(portName string, baudrate uint) {
	options := serial.OpenOptions{
		PortName:        portName,
		BaudRate:        baudrate,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}
	// Open the port.
	test, err := serial.Open(options)
	myport = test
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}
}

func ClosePort() {
	defer myport.Close()
}

func SendPelco(payload []byte) {
	n, err := myport.Write(payload)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	fmt.Println("Wrote", n, "bytes.")
}
