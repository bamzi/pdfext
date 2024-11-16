package endian

import (
	_e "encoding/binary"
	_f "unsafe"
)

func IsBig() bool { return _b }
func init() {
	const _fe = int(_f.Sizeof(0))
	_c := 1
	_g := (*[_fe]byte)(_f.Pointer(&_c))
	if _g[0] == 0 {
		_b = true
		ByteOrder = _e.BigEndian
	} else {
		ByteOrder = _e.LittleEndian
	}
}
func IsLittle() bool { return !_b }

var (
	ByteOrder _e.ByteOrder
	_b        bool
)
