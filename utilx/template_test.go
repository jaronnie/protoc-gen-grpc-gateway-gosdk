package utilx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	Duplicate []string
}

var TestTpc = `
hello{{range $v := .Duplicate | uniq}}
{{$v | firstLower}}{{end}}
`

func TestParseTemplate(t *testing.T) {
	t.Run("run register", func(t *testing.T) {
		template, err := ParseTemplate(TestData{
			Duplicate: []string{"abc", "abc", "ABC", "Abc", "ABc"},
		}, []byte(TestTpc))
		assert.Nil(t, err)
		assert.Equal(t, "\nhello\nabc\naBC\nabc\naBc\n", string(template))
	})
}

type TestAnyData struct {
	Test []string
	All  []string
}

var TestAny = `{{range $v := .Test}}
{{if has $v $.All}} 111 {{end}}
{{end}}
`

func TestParseTemplateWithAny(t *testing.T) {
	t.Run("run register", func(t *testing.T) {
		template, err := ParseTemplate(TestAnyData{
			All:  []string{"abc", "abc", "ABC", "Abc", "ABc"},
			Test: []string{"abc"},
		}, []byte(TestAny))

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(template))
	})
}
