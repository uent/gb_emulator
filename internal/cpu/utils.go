package cpu

func jointBytesToUInt16(h byte, l byte) uint16 {
	return uint16(h)<<8 | uint16(l)
}

// return (high, low)
func splitUInt16ToBytes(value uint16) (byte, byte) {
	low := byte(value & 0x00FF)
	high := byte((value >> 8) & 0xFF)

	return high, low
}
