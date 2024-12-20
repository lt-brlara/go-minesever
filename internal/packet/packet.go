package packet

import (
	"bytes"
	"errors"

	"github.com/blara/go-mineserver/internal/state"
)

const (
	STATUS_PACKET_ID        byte = 0x00
	PING_PACKET_ID          byte = 0x01
	CUSTOM_REPORT_PACKET_ID byte = 0x7A
)

var (
	ErrPacketNotHandled = errors.New("Packet does not have matching struct")
)

// A Request is the generic representation of serverbound information.
type Request interface{}

// A Response is the generic interface for all clientbound information.
//
// Serialize converts the struct to byte format represented by a bytes.Buffer.
type Response interface {
	Serialize() (bytes.Buffer, error)
}

// RequestFactory returns a Request with the proper struct fields matching
// the protocol specification based on bytes read from data.
func RequestFactory(data *bytes.Buffer, session *state.Session) (Request, error) {

	length, err := readVarInt(data)
	if err != nil {
		return nil, err
	}

	packetID, err := readVarInt(data)
	if err != nil {
		return nil, err
	}
	packetIDByte := byte(packetID)

	switch packetIDByte {
	case STATUS_PACKET_ID:
		return StatusPacketFactory(packetIDByte, length, data, session)
	case PING_PACKET_ID:
		return NewPingRequest(data)
	}

	return nil, ErrPacketNotHandled
}

// StatusPacketFactory returns the correct Request based on the criteria of
// different types of status-related packets.
func StatusPacketFactory(id byte, length int32, data *bytes.Buffer, session *state.Session) (Request, error) {

	switch session.State {
	case state.StateNull:
		return NewHandshakeRequest(data)
	case state.StateStatus:
		return NewStatusRequest(data)
	}

	return nil, ErrPacketNotHandled
}
