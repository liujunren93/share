module github.com/liujunren93/share

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/golang/protobuf v1.4.2
	github.com/liujunren93/share_utils v0.0.0-20201207084023-8eec376a644e
	github.com/opentracing/opentracing-go v1.2.0
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
	github.com/sirupsen/logrus v1.6.0
	google.golang.org/grpc v1.31.0

)
