module github.com/liujunren93/share

go 1.15

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	google.golang.org/grpc v1.31.0 => google.golang.org/grpc v1.26.0
)

require (
	github.com/containerd/containerd v1.3.0
	github.com/envoyproxy/protoc-gen-validate v0.4.1
	github.com/golang/protobuf v1.4.3
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/liujunren93/share_utils v0.0.0-20201226110723-fb0ad234bbd8
	github.com/opentracing/opentracing-go v1.2.0
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v0.0.3 // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	google.golang.org/grpc v1.31.0
)
