package cSystemInfo

const (
// Add Constants here
)

type SystemInfoIF interface {
	GetIP()
}

func GetIP(message SystemInfoIF) {
	message.GetIP()
}
