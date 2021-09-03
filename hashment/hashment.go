package hashment

const (
	IntBits32=32
	IntBits64=64
)

func LeftRotateInt32(n int32, d uint8) int32{
	return (n << d) | (n >> (IntBits32 - d))
}

func RighRotateInt32(n int32, d uint8) int32 {
	return (n >> d) | (n << (IntBits32 - d))
}

func RighRotateUint32(n uint32, d uint8) uint32 {
	return (n >> d) | (n << (IntBits32 - d))
}

func RighRotateInt64(n int64, d uint8) int64 {
	return (n >> d) | (n << (IntBits64 - d))
}

func HashCodeSimple(i int) int {
	return int(RighRotateInt32(int32(i), 32)) + i
}

func HashCodeOther(i int) int {
	return int(RighRotateInt32(int32(i), 31)) ^ i
}

func HashCodePolynomial(s string) uint32 {
	var h uint32
	for i := 0; i < len(s); i++ {
		h = (h << 5) | RighRotateUint32(h, 27)
		h += uint32(s[i])
	}

	return h
}

func CompressHashCodeDivision(v uint32, l int) int {
	return int(v) % l
}