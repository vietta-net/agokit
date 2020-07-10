package helper

import "encoding/json"

func Convert(from interface{}, to interface{}) (err error) {
	jData, err := json.Marshal(from)
	json.Unmarshal(jData, &to)
	return err
}
