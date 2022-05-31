package client

import (
	"fmt"

	"github.com/liujunren93/share/core/resolver/endpoint"
)

func BuildDirectTarget(namespace, srvName string) string {

	return fmt.Sprintf("%s:///%s/%s", endpoint.Name, namespace, srvName)
}
