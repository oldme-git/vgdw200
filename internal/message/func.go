package message

// GetChecksum 获取校检字
func GetChecksum(data []byte) byte {
	checksum := byte(0)
	for _, b := range data {
		checksum ^= b
	}
	return checksum
}

// DataBytesBig16 获取uint16的大端序数组byte
func DataBytesBig16(l uint16) [2]byte {
	return [2]byte{
		byte(l),
		byte(l >> 8),
	}
}

// DataBytesBig32 获取uint32的大端序数组byte
func DataBytesBig32(l uint32) [4]byte {
	return [4]byte{
		byte(l),
		byte(l >> 8),
		byte(l >> 16),
		byte(l >> 24),
	}
}
