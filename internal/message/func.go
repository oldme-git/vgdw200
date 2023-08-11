package message

// GetChecksum 获取校检字
func GetChecksum(data []byte) byte {
	checksum := byte(0)
	for _, b := range data {
		checksum ^= b
	}
	return checksum
}

// DataBytesBig 获取uint16的大端序数组byte
func DataBytesBig(l uint16) [2]byte {
	// 获取高位字节
	highByte := byte(l >> 8)
	lowByte := byte(l)
	return [2]byte{lowByte, highByte}
}
