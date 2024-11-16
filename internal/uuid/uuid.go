package uuid

import (
	_c "crypto/rand"
	_cb "encoding/hex"
	_ac "io"
)

var _gg UUID

type UUID [16]byte

func NewUUID() (UUID, error) {
	var uuid UUID
	_, _ad := _ac.ReadFull(_af, uuid[:])
	if _ad != nil {
		return _gg, _ad
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return uuid, nil
}
func _ec(_ecd []byte, _b UUID) {
	_cb.Encode(_ecd, _b[:4])
	_ecd[8] = '-'
	_cb.Encode(_ecd[9:13], _b[4:6])
	_ecd[13] = '-'
	_cb.Encode(_ecd[14:18], _b[6:8])
	_ecd[18] = '-'
	_cb.Encode(_ecd[19:23], _b[8:10])
	_ecd[23] = '-'
	_cb.Encode(_ecd[24:], _b[10:])
}
func (_f UUID) String() string { var _ea [36]byte; _ec(_ea[:], _f); return string(_ea[:]) }

var Nil = _gg
var _af = _c.Reader

func MustUUID() UUID {
	uuid, _e := NewUUID()
	if _e != nil {
		panic(_e)
	}
	return uuid
}
