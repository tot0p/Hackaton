package utils

func IsIterable(object any) bool {
	switch object.(type) {
	case string, int, bool, float64, float32, complex64, complex128, uint, uint8, uint16, uint32, uint64, int8, int16, int32, int64:
		return false
	default:
		return true
	}
}
