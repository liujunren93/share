package status

type status interface {
	Code() int32
	Msg() string
}

