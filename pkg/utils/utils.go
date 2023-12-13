package utils

import (
	"Dictionary/models/rows"
	"encoding/json"
)

func MarshalValue(data *rows.Rows) ([]byte, error) {
	// Création d'une carte (map)
	myMap := map[string][]byte{
		"desc": data.Desc,
		"time": data.Date,
	}

	return json.Marshal(myMap)
}

func UnmarshalValue(bytes []byte) (*rows.Rows, error) {
	var um map[string][]byte
	err := json.Unmarshal(bytes, &um)
	if err != nil {
		return &rows.Rows{}, err
	}
	return &rows.Rows{Desc: um["desc"], Date: um["time"]}, nil
}
