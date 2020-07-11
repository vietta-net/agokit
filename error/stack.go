package error

import "encoding/json"

type Stack struct {
	Callers []uintptr
}

func (t Stack) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Stack) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &t)
}