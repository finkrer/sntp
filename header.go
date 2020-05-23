package main

import (
	"encoding/binary"
	"errors"
)

type sntpHeader struct {
	LI,
	VN,
	Mode,
	Stratum,
	Poll,
	Precision uint8
	Delay,
	Dispersion,
	ReferenceID uint32
	ReferenceTimestamp,
	OriginateTimestamp,
	ReceiveTimestamp,
	TransmitTimestamp uint64
}

func parseHeader(header []byte) (sntpHeader, error) {
	if len(header) < 48 {
		return sntpHeader{}, errors.New("Header was too short")
	}

	vn := readBits(header[0], 2, 3)
	mode := readBits(header[0], 5, 3)
	poll := uint8(header[2])
	trans := binary.BigEndian.Uint64(header[40:48])

	return sntpHeader{
		VN: vn, Mode: mode, Poll: poll, TransmitTimestamp: trans,
	}, nil
}

func (h *sntpHeader) toBinary() []byte {
	result := make([]byte, 48)
	writeBits(&result[0], 0, 2, h.LI)
	writeBits(&result[0], 2, 3, h.VN)
	writeBits(&result[0], 5, 3, h.Mode)
	result[1] = byte(h.Stratum)
	result[2] = byte(h.Poll)
	result[3] = byte(h.Precision)
	binary.BigEndian.PutUint32(result[4:8], h.Delay)
	binary.BigEndian.PutUint32(result[8:12], h.Dispersion)
	binary.BigEndian.PutUint32(result[12:16], h.ReferenceID)
	binary.BigEndian.PutUint64(result[16:24], h.ReferenceTimestamp)
	binary.BigEndian.PutUint64(result[24:32], h.OriginateTimestamp)
	binary.BigEndian.PutUint64(result[32:40], h.ReceiveTimestamp)
	binary.BigEndian.PutUint64(result[40:48], h.TransmitTimestamp)

	return result
}
