package config_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/geraldkohn/elian/config"
)

func TestLicense(t *testing.T) {
	l := &config.License{
		LicenseCode: "test",
		LicenseOwner: "admin",
	}
	b, err := json.Marshal(l)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

func TestLicenseUnMarshal(t *testing.T) {
	l := []config.License{}
	json.Unmarshal(config.LicenseData, &l)
	fmt.Println(l[0].LicenseCode)
	fmt.Println(l[0].LicenseOwner)
	fmt.Println(l[1].LicenseCode)
	fmt.Println(l[1].LicenseOwner)
}