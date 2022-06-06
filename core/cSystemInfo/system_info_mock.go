/*
System Info retrieves information in the SYC format

RESTRICTIONS:
    None

NOTES:
    None
*/
package core

import (
	"net"
)

type SystemInfoMock struct {
	InternalIP net.Addr
}

// SystemInfoMock.GetIP
func (si SystemInfoMock) GetIP() net.Addr {
	return nil
}
