package pelcod

//Speed from 0x01-3F
func PanLeft(address byte, speed byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3PanLeft, byte4PanLeft, speed, 0x00, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

func PanRight(address byte, speed byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3PanRight, byte4PanRight, speed, 0x00, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

func TiltUp(address byte, speed byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3TiltUp, byte4TiltUp, 0x00, speed, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

func TiltDown(address byte, speed byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3TiltDown, byte4TiltDown, 0x00, speed, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

func ZoomTele(address byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3ZoomTele, byte4ZoomTele, 0x00, 0x00, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

func ZoomWide(address byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3ZoomTele, byte4ZoomTele, 0x00, 0x00, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

func FocusFar(address byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3FocusFar, byte4FocusFar, 0x00, 0x00, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

func FocusNear(address byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3FocusNear, byte4FocusNear, 0x00, 0x00, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

func PTZFStop(address byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3PTZFStop, byte4PTZFStop, 0x00, 0x00, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}
