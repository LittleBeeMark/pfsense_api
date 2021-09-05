package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"pfsense/pf_client"
)

func main() {
	cl := pf_client. Cl{
		BaseUrl:     "https://192.168.252.183/api/v1/",
		ClientToken: "pfsense",
		ClientID:    "admin",
	}

	reqBody := struct {
		ClientID    string `json:"client-id"`
		ClientToken string `json:"client-token"`
	}{
		ClientID:    cl.ClientID,
		ClientToken: cl.ClientToken,
	}

	reqB, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	err = cl.Request(context.Background(), http.MethodGet, "/system/certificate", "", bytes.NewReader(reqB), nil)
	if err != nil {
		fmt.Println("req err:", err)
	}
}
