package ws

import (
	"encoding/json"
)

type InputEnvelop struct {
	Message string          `json:"message"`
	Room    string          `json:"room"`
	Data    json.RawMessage `json:"data"`
}

func (e *InputEnvelop) Decode(data []byte) (*InputEnvelop, error) {
	err := json.Unmarshal(data, e)
	return e, err
}

type OutputEnvelop struct {
	Data     json.RawMessage `json:"data"`
	Received int64           `json:"received"`
}

func (e *OutputEnvelop) Encode() ([]byte, error) {
	return json.Marshal(e)
}
