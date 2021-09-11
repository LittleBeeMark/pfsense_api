package main

import (
	"pfsense/cert"
)

func main() {
	param := &cert.EditCertParam{}

	cert.EditPfSenseCert(param)

}
