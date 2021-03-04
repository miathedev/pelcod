package pelcod

import (
	"fmt"
	"io"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

var port io.ReadWriteCloser

func openSerialPort() {
	options := serial.OpenOptions{
		PortName:        "/dev/tty.usbserial-A8008HlV",
		BaudRate:        19200,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}
	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}
	port
}

func closePort() {
	defer port.Close()
}

func sendPelco(payload []byte) {
	n, err := port.Write(payload)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	fmt.Println("Wrote", n, "bytes.")
}
