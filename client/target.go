package client

import (
	"fmt"

	"github.com/liujunren93/share/core/resolver/endpoint"
)

// namespace, srvName string
func defaultDirectTarget(args ...string) string {
	return fmt.Sprintf("%s:///%s/%s", endpoint.Name, args[0], args[1])
}
