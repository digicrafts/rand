package rand

import (
	"crypto/rand"
	"encoding/binary"
)

type CryptoRandSource struct{}

func NewCryptoRandSource() CryptoRandSource {
	return CryptoRandSource{}
}

func (s CryptoRandSource) Uint64() (value uint64) {
	binary.Read(rand.Reader, binary.BigEndian, &value)
	return value
}

func (s CryptoRandSource) Int63() int64 {
	// Call Uint64
	// make sure the most significant bit is a always a 0
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s CryptoRandSource) Seed(_ int64) {}
