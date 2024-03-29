package utils

import (
	"math/rand"
	"net"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetUuidV3(name string) string {
	v1, _ := uuid.NewV4()
	variant := uuid.NewV3(v1, name)
	//all := strings.ReplaceAll(variant.String(), "-", "")
	return variant.String()
}

func GetIntranetIp() (net.IP, error) {
	conn, err := net.Dial("udp", "8.0.0.1:8")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
	// addrs, err := net.InterfaceAddrs()

	// if err != nil {
	// 	return nil,err

	// }

	// for _, address := range addrs {
	// 	// 检查ip地址判断是否回环地址
	// 	if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	// 		if ipnet.IP.To4() != nil {
	// 			return ipnet.IP,nil
	// 		}

	// 	}
	// }
	// return nil, nil
}

func RandString(length int) string {
	baseStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63n(9999)))
	var randStr []byte
	buf := []byte(baseStr)
	for i := 0; i < length; i++ {
		randStr = append(randStr, buf[r.Intn(len(baseStr))])
	}
	return string(randStr)
}

func RandInt(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63n(9999)))
	return r.Intn(max)
}

func RandIntRange(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63n(9999)))
	return r.Intn(max-min) + min
}
