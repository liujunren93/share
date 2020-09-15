package status

type status interface {
	GetCode() int32
	GetMsg() string
}

