package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfPlayFreeGacha represents the MSG_MHF_PLAY_FREE_GACHA
type MsgMhfPlayFreeGacha struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfPlayFreeGacha) Opcode() network.PacketID {
	return network.MSG_MHF_PLAY_FREE_GACHA
}

// Parse parses the packet from binary
func (m *MsgMhfPlayFreeGacha) Parse(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfPlayFreeGacha) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}