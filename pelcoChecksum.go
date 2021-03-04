package pelcod

import (
	"fmt"
	"strconv"
)

func calcPelcoChecksum(packet *[7]byte) byte {
	checksum := (uint(packet[1]) + uint(packet[2]) + uint(packet[3]) + uint(packet[4]) + uint(packet[5])) % 256
	packet[6] = byte(checksum)
	fmt.Println("Checksum: " + strconv.Itoa(int(checksum)))
	return byte(checksum)
}
