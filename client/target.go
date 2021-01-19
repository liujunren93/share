package client

import (
	"fmt"
	"github.com/liujunren93/share/core/resolver/endpoint"
)

const DirectScheme = "dns"

func BuildDirectTarget(serverName string) string {

	return fmt.Sprintf("%s:///%s", endpoint.Name, serverName)
}
