package utilx

import (
	"bytes"
	"encoding/json"
)

func BeautifyJson(v interface{}) (string, error) {
	uglyBody, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	err = json.Indent(&out, uglyBody, "", "\t")
	return out.String(), nil
}
