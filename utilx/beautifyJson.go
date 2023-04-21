package utilx

import (
	"bytes"
	"encoding/json"
)

func BeautifyJSON(v interface{}) (string, error) {
	uglyBody, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	err = json.Indent(&out, uglyBody, "", "\t")
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
