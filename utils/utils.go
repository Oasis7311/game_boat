package utils

func PtrI64(val int64) *int64 {
	res := new(int64)
	*res = val
	return res
}

func PtrInt(val int) *int {
	res := new(int)
	*res = val
	return res
}

func PtrI32(val int32) *int32 {
	res := new(int32)
	*res = val
	return res
}

func PtrUint64(val uint64) *uint64 {
	res := new(uint64)
	*res = val
	return res
}

func PtrUint(val uint) *uint {
	res := new(uint)
	*res = val
	return res
}

func PtrStr(val string) *string {
	res := new(string)
	*res = val
	return res
}
