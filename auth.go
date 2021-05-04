package netapp

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	basePath = "/api/ontap/"
	method   = "GET"
	user     = os.Getenv("USER")
	password = os.Getenv("PASSWORD")
	server   = os.Getenv("SERVER")
)

type NewClient struct {
	UserName  string
	Password  string
	Host      string
	VerifySSL bool
	SSL       bool
	TimeOut   time.Duration
	URL       string
	Debug     bool
	Client    *http.Client
}

// Instantiate new client
func client(host, username, password string, ssl bool, timeout time.Duration) *NewClient {
	url := "https://" + host + basePath
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return &NewClient{
		UserName:  username,
		Password:  password,
		Host:      host,
		VerifySSL: ssl,
		SSL:       true,
		TimeOut:   timeout,
		URL:       url,
		Debug:     false,
		Client:    client,
	}
}

//Auth creates a Ontap/Ocum client
func Auth(server, user, password string) *NewClient {
	client := client(server, user, password, true, 60*time.Second)
	return client
}

func clientV2(host, username, password string, ssl bool, timeout time.Duration) *NewClient {

	url := "https://" + host
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return &NewClient{
		UserName:  username,
		Password:  password,
		Host:      host,
		VerifySSL: ssl,
		SSL:       true,
		TimeOut:   timeout,
		URL:       url,
		Debug:     false,
		Client:    client,
	}
}

func AuthV2(server, user, password string) *NewClient {
	client := clientV2(server, user, password, true, 60*time.Second)
	return client
}

func getResponseBody(query string) ([]byte, error) {
	client := AuthV2(server, user, password)
	newurl := client.URL + query
	req, err := http.NewRequest(method, newurl, nil)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	req.SetBasicAuth(client.UserName, client.Password)
	resp, err := client.Client.Do(req)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return bodyText, nil
}
