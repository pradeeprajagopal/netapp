package netapp

import (
	"crypto/tls"
	"net/http"
	"time"
)

var (
	basePath = "/api/ontap/"
	method   = "GET"
	server string
	user string
	password string
)

type NewClient struct {
	UserName   string
	Password   string
	Host       string
	VerifySSL  bool
	SSL        bool
	TimeOut    time.Duration
	URL        string
	Debug      bool
	Client     *http.Client
}

// Instantiate new client
func client(host, username, password string, ssl bool, timeout time.Duration) *NewClient {
	url := "https://" + host + basePath
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return &NewClient{
		UserName:   username,
		Password:   password,
		Host:       host,
		VerifySSL:  ssl,
		SSL:        true,
		TimeOut:    timeout,
		URL:        url,
		Debug:      false,
		Client:     client,
	}
}

//Auth creates a Ontap/Ocum client
func Auth(server, user, password string) *NewClient {
	client := client(server, user, password, true, 60*time.Second)
	return client
}
