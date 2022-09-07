package json

import (
	js "encoding/json"

	"google.golang.org/grpc/encoding"
)

const Name = "share_json"

type Codes struct{}

func init() {
	encoding.RegisterCodec(Codes{})
}

//c codes

func (Codes) Marshal(v interface{}) ([]byte, error) {
	if vv, ok := v.([]byte); ok {
		return vv, nil
	}
	return js.Marshal(v)
}
func (Codes) Unmarshal(data []byte, v interface{}) error {
	if len(data) == 0 {
		return nil
	}
	return js.Unmarshal(data, v)
}
func (Codes) Name() string {
	return Name
}

func (Codes) String() string {
	return Name
}
