package pelcod

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

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

//0x00-7F
func SetPreset(address byte, preset byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3SetPreset, byte4SetPreset, 0x00, preset, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

//0x00-8B
func GotoPreset(address byte, preset byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3GotoPreset, byte4GotoPreset, 0x00, preset, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

//0x00-7F
func ClearPreset(address byte, preset byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3ClearPreset, byte4ClearPreset, 0x00, preset, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

//0x01-06
func AuxOn(address byte, aux byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3AuxOn, byte4AuxOn, 0x00, aux, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

//0x01-06
func AuxOff(address byte, aux byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3AuxOff, byte4AuxOff, 0x00, aux, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

func CameraOn(address byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3CameraOn, byte4CameraOn, 0x00, 0x00, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

func CameraOff(address byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3CameraOff, byte4CameraOff, 0x00, 0x00, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

//0x00-03
func SetZoomSpeed(address byte, speed byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3SetZoomSpeed, byte4SetZoomSpeed, 0x00, speed, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

//0x00-03
func SetFocusSpeed(address byte, speed byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3SetFocusSpeed, byte4SetFocusSpeed, 0x00, speed, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

func ResetToDefaults(address byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3ResetToDefaults, byte4ResetToDefaults, 0x00, 0x00, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

//0x00 -> Auto
//0x01 -> On
//0x02 -> Off
func AutoFocusOnOff(address byte, focus byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3AutoFocusAutoOnOff, byte4AutoFocusAutoOnOff, 0x00, focus, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

//0x00 -> Auto
//0x01 -> On
//0x02 -> Off
func AutoIrisOnOff(address byte, iris byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3AutoIrisAutoOnOff, byte4AutoIrisAutoOnOff, 0x00, iris, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

//0x00 -> Auto
//0x01 -> On
//0x02 -> Off
func AGCOnOff(address byte, agc byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3AGCAutoOnOff, byte4AGCAutoOnOff, 0x00, agc, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

//0x00 -> /
//0x01 -> On
//0x02 -> Off
func BacklightCompensation(address byte, backlightCompensation byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3BacklightCompensationOnOff, byte4BacklightCompensationOnOff, 0x00, backlightCompensation, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

//0x00 -> /
//0x01 -> On
//0x02 -> Off
func AutoWhiteBalanceOnOff(address byte, autoWhiteBalance byte) [7]byte {
	payload := [7]byte{byte1SyncByte, address, byte3AutoWhiteBalanceOnOff, byte4AutoWhiteBalanceOnOff, 0x00, autoWhiteBalance, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

func getMsbLsb(input uint16) (byte, byte) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, input)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	lsb := buf.Bytes()[0]
	msb := buf.Bytes()[1]
	return msb, lsb
}

func SetPanPosition(address byte, pan uint16) [7]byte {
	msb, lsb := getMsbLsb(pan)
	payload := [7]byte{byte1SyncByte, address, byte3SetPanPosition, byte4SetPanPosition, msb, lsb, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

func SetTiltPosition(address byte, tilt uint16) [7]byte {
	msb, lsb := getMsbLsb(tilt)
	payload := [7]byte{byte1SyncByte, address, byte3SetTiltPosition, byte4SetTiltPosition, msb, lsb, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

func SetZoomPosition(address byte, zoom uint16) [7]byte {
	msb, lsb := getMsbLsb(zoom)
	payload := [7]byte{byte1SyncByte, address, byte3SetZoomPosition, byte4SetZoomPosition, msb, lsb, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}

func SetFocusPosition(address byte, focus uint16) [7]byte {
	msb, lsb := getMsbLsb(focus)
	payload := [7]byte{byte1SyncByte, address, byte3SetFocusPosition, byte4SetFocusPosition, msb, lsb, 0x00}
	calcPelcoChecksum(&payload)
	return payload
}
