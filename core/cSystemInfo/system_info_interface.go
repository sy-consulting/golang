package core

type SystemInfoIF interface {
	GetIP()
}

func GetIP(siIF SystemInfoIF) {
	siIF.GetIP()
}
