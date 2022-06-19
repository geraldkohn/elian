package config

type License struct {
	LicenseCode  string `json:"license_code"`
	LicenseOwner string `json:"license_owner"`
}

var LicenseData = []byte(`[
{"license_code":"test","license_owner":"admin"},
{"license_code":"test1","license_owner":"admin1"}
]`)
