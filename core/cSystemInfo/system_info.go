/*
System Info retrieves information in the SYC format

RESTRICTIONS:
    None

NOTES:
    None
*/
package cSystemInfo

import (
	"net"
)

type SystemInfo struct {
	InternalIP net.Addr
}

// SystemInfo.GetIP
func (si SystemInfo) GetIP() net.Addr {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	return conn.LocalAddr().(*net.UDPAddr)
}
