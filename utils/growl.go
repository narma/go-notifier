package utils

import (
	"fmt"
	"net"
	"time"
)

const growlPort = 23053

func HasGrowl() bool {
	dialer := net.Dialer{
		Timeout:   time.Millisecond * 100,
		KeepAlive: 0,
	}
	conn, err := dialer.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", growlPort))
	if err != nil {
		return false
	}

	defer conn.Close()
	return true
}
