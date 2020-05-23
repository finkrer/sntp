package main

import (
	"math"
	"time"
)

var ntpEpoch = time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)

func (h *sntpHeader) getResponse() *sntpHeader {
	li := uint8(0)
	vn := h.VN
	mode := h.Mode + 1
	stratum := uint8(1)
	poll := h.Poll
	precision := uint8(247)
	delay := uint32(0)
	dispersion := uint32(0)
	refID := uint32(0b01001100_01001111_01000011_01001100)

	orig := h.TransmitTimestamp
	trans := getTime()
	ref := trans
	recv := trans

	return &sntpHeader{
		li, vn, mode, stratum, poll, precision, delay, dispersion, refID, ref, orig, recv, trans,
	}
}

func getTime() uint64 {
	delayed := time.Now().UTC().Add(time.Duration(delay) * time.Second)
	t := delayed.Sub(ntpEpoch)
	s := t.Seconds()
	ns := t.Nanoseconds() - int64(s)*1E9
	frac := uint64(ns) * uint64(math.Pow(2, 32)/1E9)
	return (uint64(s) << 32) + frac
}
