package helper

import (
	"context"
	"errors"
	"net"
	"strings"
)

func IsEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func GetHostName() string {
	return GetLocalIP().String()
}

func GetLocalIP() net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP
			}
		}
	}
	return nil
}

func ExtractFromContext(ctx context.Context, field string) (string, error) {
	if ctx == nil {
		return "", errors.New("Context is empty")
	}

	value := ctx.Value(field)
	if value == nil {
		return "", errors.New("Value is nil")
	}
	return value.(string), nil
}
