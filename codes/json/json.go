package json

import (
	js "encoding/json"

	"google.golang.org/grpc/encoding"
)

const Name = "json"

type Codes struct{}

func init() {
	encoding.RegisterCodec(Codes{})
}

//c codes

func (Codes) Marshal(v interface{}) ([]byte, error) {
	return js.Marshal(v)
}
func (Codes) Unmarshal(data []byte, v interface{}) error {
	return js.Unmarshal(data, v)
}
func (Codes) Name() string {
	return Name
}

func (Codes) String() string {
	return Name
}
