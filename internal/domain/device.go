package domain

type Device struct {
	ID        string
	Name      string
	Cert      string
	CreatedAt string
}

type DeviceAuthResult struct {
	IsValid bool
	Reason  string
}
