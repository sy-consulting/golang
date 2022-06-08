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
	"strings"
)

type SystemInfo struct {
	Owner string
}

// SystemInfo.GetIP
func (SystemInfo) GetIP() (x string) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
