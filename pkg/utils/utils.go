package utils

import (
	"Dictionary/models/rows"
	"encoding/json"
)

// MarshalValue converts a rows.Rows object into a JSON byte array.
// It creates a map containing the desc and time values from the rows.Rows object,
// then uses json.Marshal to encode this map into JSON bytes.
func MarshalValue(data *rows.Rows) ([]byte, error) {
	myMap := map[string][]byte{
		"desc": data.Desc,
		"time": data.Date,
	}

	return json.Marshal(myMap)
}

// UnmarshalValue converts a JSON byte array into a rows.Rows object.
// It decodes the JSON byte array into a map and extracts the desc and time values.
// Using these values, it creates and returns a new rows.Rows object.
func UnmarshalValue(bytes []byte) (*rows.Rows, error) {
	var myMap map[string][]byte
	err := json.Unmarshal(bytes, &myMap)
	if err != nil {
		return &rows.Rows{}, err
	}
	return &rows.Rows{Desc: myMap["desc"], Date: myMap["time"]}, nil
}
