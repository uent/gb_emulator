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

func calculateHalfFlagAdd(oldValue byte, addedValue byte) bool {
	return ((oldValue & 0x0F) + (addedValue & 0x0F)) > 0x0F
}

func calculateHalfFlagSubtract(oldValue byte, subtractedValue byte) bool {
	return (oldValue & 0x0F) < (subtractedValue & 0x0F)
}

func calculateHalfFlagIncrement(oldValue byte) bool {
	return (oldValue&0x0F)+1 > 0x0F
}

func bool2u8(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}
