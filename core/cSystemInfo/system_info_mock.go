/*
System Info retrieves information in the SYC format

RESTRICTIONS:
    None

NOTES:
    None
*/
package cSystemInfo

type SystemInfoMock struct {
	Owner string
}

// SystemInfoMock.GetIP
func (SystemInfoMock) GetIP() string {
	return "0.0.0.0"
}
