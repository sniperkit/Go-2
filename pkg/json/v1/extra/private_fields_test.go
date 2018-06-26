package extra

import (
	"testing"

	// external
	"github.com/stretchr/testify/require"

	// internal
	"github.com/sniperkit/snk.golang.json/pkg/json/v1"
)

func Test_private_fields(t *testing.T) {
	type TestObject struct {
		field1 string
	}
	SupportPrivateFields()
	should := require.New(t)
	obj := TestObject{}
	should.Nil(jsoniter.UnmarshalFromString(`{"field1":"Hello"}`, &obj))
	should.Equal("Hello", obj.field1)
}
