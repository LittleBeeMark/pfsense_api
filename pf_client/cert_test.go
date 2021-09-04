package pf_client

import (
	"testing"
)

func TestImportCert(t *testing.T) {
	ImportCert()
	//	fmt.Println(base64.StdEncoding.EncodeToString([]byte(mCert)))
}

func TestReadCerts(t *testing.T) {
	ReadCerts()
}
