package packet

import (
	"bytes"
	"encoding/json"

	"github.com/blara/go-mineserver/internal/log"
)

// A StatusRequest represents all fields in a serverbound status_request.
//
// See the reference documentation on status_request for more information:
// https://wiki.vg/Protocol#Status
type StatusRequest struct{}

// NewStatusRequest returns a StatusRequest with all fields parsed.
func NewStatusRequest(data *bytes.Buffer) (*StatusRequest, error) {
	return &StatusRequest{}, nil
}

// A StatusResponse represents all fields in a serverbound status_response.
//
// See the reference documentation on status_response for more information:
// https://wiki.vg/Protocol#Status
type StatusResponse struct {
	Version            `json:"version"`
	Players            `json:"players"`
	Description        `json:"description"`
	Favicon            string `json:"favicon"`
	EnforcesSecureChat bool   `json:"enforcesSecureChat"`
}

func NewStatusReponse() *StatusResponse {
	resp := &StatusResponse{
		Version: Version{
			Name:     "1.21.1",
			Protocol: 767,
		},
		Players: Players{
			Max:    20,
			Online: 0,
		},
		Description: Description{
			Text: "Test server!",
		},
		EnforcesSecureChat: false,
	}

	return resp
}

func (r *StatusResponse) Serialize() (bytes.Buffer, error) {

	bytePayload, err := json.Marshal(r)
	if err != nil {
		log.Error("Could not marshal data", "error", err)
	}

	var buf bytes.Buffer
	WriteVarInt(&buf, int32(STATUS_PACKET_ID))
	WriteVarInt(&buf, int32(len(bytePayload)))
	buf.Write(bytePayload)

	var resp bytes.Buffer
	WriteVarInt(&resp, int32(buf.Len()))
	resp.Write(buf.Bytes())

	return resp, err
}

type Version struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type Players struct {
	Max    int      `json:"max"`
	Online int      `json:"online"`
	Sample []Player `json:"sample"`
}

type Player struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type Description struct {
	Text string `json:"text"`
}