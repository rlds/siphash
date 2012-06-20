package siphash

import "testing"

func TestNew(t *testing.T) {
	k := []byte{0x00,0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x09,0x0a,0x0b,0x0c,0x0d,0x0e,0x0f}
	m := []byte{0x00,0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x09,0x0a,0x0b,0x0c,0x0d,0x0e}
	result := uint64(0xa129ca6149be45e5)
	h := New(k)
	h.Write(m)
	if sum := h.Sum64(); sum != result {
		t.Errorf("expected %x, got %x", result, sum)
	}
}
