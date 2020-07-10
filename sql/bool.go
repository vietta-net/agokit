package sql

func BoolToUint32(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}

func Uint32ToBool(i uint32) bool {
	if i > 0 {
		return true
	}
	return false
}
