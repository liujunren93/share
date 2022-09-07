package etcd

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/liujunren93/share/core/registry"
)

var r *etcdRegistry

func init() {
	// /registry/shareLife/share_app_rbac
	etcdRegistry, err := NewRegistry(registry.WithAddrs("node1:2379",
		"node1:3379",
		"node1:4379"), registry.WithPrefix("registry"))
	if err != nil {
		panic(err)
	}
	r = etcdRegistry
}
func TestNewRegistry(t *testing.T) {
	go func() {
		r.Watch(context.TODO(), "aaaaaa/test")
	}()
	for {

		ser, err := registry.RegistryInstance.GetService(context.TODO(), "aaaaaa/test")

		fmt.Println(ser, err)
		time.Sleep(time.Second * 3)
	}

}
