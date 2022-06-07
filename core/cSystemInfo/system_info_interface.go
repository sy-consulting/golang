package cSystemInfo

type SystemInfoIF interface {
	GetIP() string
}

func GetIP(siIF SystemInfoIF) (x string) {
	return siIF.GetIP()
}
