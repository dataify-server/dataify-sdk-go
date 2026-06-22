package dataify

import (
	"encoding/json"
	"strings"
)

type RawResponse struct{ raw json.RawMessage }

func NewRawResponse(body []byte) (*RawResponse, error) {
	raw, err := normalizeJSON(body)
	if err != nil {
		return nil, err
	}
	return &RawResponse{raw: raw}, nil
}

func (r *RawResponse) Bytes() []byte {
	if r == nil {
		return nil
	}
	return append([]byte(nil), r.raw...)
}
func (r *RawResponse) String() string {
	if r == nil {
		return ""
	}
	return string(r.raw)
}
func (r *RawResponse) Decode(v any) error {
	if r == nil {
		return nil
	}
	return json.Unmarshal(r.raw, v)
}
func (r *RawResponse) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	return r.raw.MarshalJSON()
}

func normalizeJSON(body []byte) (json.RawMessage, error) {
	var value any
	if err := json.Unmarshal(body, &value); err != nil {
		return nil, err
	}
	if text, ok := value.(string); ok {
		trimmed := strings.TrimSpace(text)
		if trimmed != "" && json.Valid([]byte(trimmed)) {
			return normalizeJSON([]byte(trimmed))
		}
	}
	normalized, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(normalized), nil
}
