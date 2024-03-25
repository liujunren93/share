package main

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else {
		fmt.Println("Home directory not found.")
		os.Exit(1)
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	config, err := rest.InClusterConfig()
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 通过 Service 名称和 Namespace 进行服务发现
	serviceName := "your-service-name"
	namespace := "your-namespace"
	clientset.CoreV1()
	service, err := clientset.CoreV1().Services(namespace).Get(serviceName, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	service.

		// 获取服务的 ClusterIP 和端口信息
		fmt.Printf("Service Name: %s\n", service.Name)
	fmt.Printf("Service ClusterIP: %s\n", service.Spec.ClusterIP)
	fmt.Printf("Service Port: %d\n", service.Spec.Ports[0].Port)

	// 可以根据服务信息执行你的逻辑，比如构建服务地址等
	// serviceAddress := fmt.Sprintf("%s:%d", service.Spec.ClusterIP, service.Spec.Ports[0].Port)
	// 使用 serviceAddress 执行相应的操作
}
