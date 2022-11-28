package utilx

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig"
)

// ParseTemplate template
func ParseTemplate(data interface{}, tplT []byte) ([]byte, error) {
	t := template.New("temp").Funcs(sprig.TxtFuncMap())
	t, err := t.Parse(string(tplT))
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}
