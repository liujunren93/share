module github.com/liujunren93/share

go 1.16

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	google.golang.org/grpc v1.31.0 => google.golang.org/grpc v1.26.0
)

require (
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	github.com/golang/protobuf v1.5.2
	github.com/liujunren93/share_utils v0.0.0-20210817150003-3ae17b275349
	github.com/opentracing/opentracing-go v1.2.0
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
	github.com/sirupsen/logrus v1.8.1
	github.com/soheilhy/cmux v0.1.5 // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/grpc v1.31.0
)
