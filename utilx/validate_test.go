package utilx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateStruct(t *testing.T) {
	type TestStruct struct {
		A int     `validate:"min=6,max=10"`
		B string  `validate:"required"`
		C float64 `validate:"required"`
	}

	t.Run("test normal", func(t *testing.T) {
		ts := TestStruct{
			A: 6,
			B: "test_B",
			C: 6.19,
		}
		err := ValidateStruct(ts)
		assert.Nil(t, err)

		err = ValidateStruct(&ts)
		assert.Nil(t, err)
	})

	t.Run("test error", func(t *testing.T) {
		ts := TestStruct{
			A: 5,
			B: "",
			C: 0,
		}
		err := ValidateStruct(ts)
		t.Logf(err.Error())
		assert.NotNil(t, err)

		err = ValidateStruct(&ts)
		t.Logf(err.Error())
		assert.NotNil(t, err)
	})
}
