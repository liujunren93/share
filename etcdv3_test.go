package share

import (
	"context"
	"fmt"
	"github.com/shareChina/share/log"
	"github.com/shareChina/share/registry"
	"github.com/shareChina/share/registry/etcd"
	"github.com/shareChina/share/utils"
	"testing"
	"time"
)

func TestNewRegistry(t *testing.T) {
	service := registry.Service{
		Name:     "aaa",
		Version:  "1",
		Metadata: nil,
		Node:     utils.GetUuidV3("share"),
	}
	etcd := etcd.NewRegistry()
	etcd.Init()
	err := etcd.Registry(&service)
	fmt.Println(err)
	time.Sleep(time.Hour)
}
func TestNewRegistry1(t *testing.T) {
	service := registry.Service{
		Name:     "aaa",
		Version:  "2",
		Metadata: nil,
		Node:     utils.GetUuidV3("share"),

	}
	etcd := etcd.NewRegistry()
	etcd.Init()
	err := etcd.Registry(&service)
	fmt.Println(err)
	time.Sleep(time.Hour)
}
func TestNewRegistryGet(t *testing.T) {

	//service := registry.Service{
	//	Name:     "aaa",
	//	Version:  "1",
	//	Metadata: nil,
	//	Node:     utils.GetUuidV3("share"),
	//
	//}
	etcd := etcd.NewRegistry()
	etcd.Init()
	getService, err := etcd.GetService("aaa")
	for _, r := range getService {
		fmt.Println(r,err)
	}

}

func TestWatch(t *testing.T)  {
	etcd := etcd.NewRegistry()
	etcd.Init()
	//timeout, _ := context.WithTimeout(context.TODO(), time.Second*5)
	service, _ := etcd.GetService("aaa")
	go etcd.Watch("aaa",context.TODO(),&service)
	for {

		fmt.Println(service)
		time.Sleep(time.Second*2)
	}



}

func TestContext(t *testing.T)  {
	log.Logger.Error(111)
}

