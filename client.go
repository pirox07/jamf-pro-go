package jamf_pro_go

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	baseURL		string
	v1Token		string
	classicToken	string
}

// V1Token is the response to the Jamf API Token request.
type V1Token struct {
	Token	string	`json:"token"`
	Expires	int64	`json:"expires"`
}

func NewConfig(url, userName, password string) (*Config, error) {
	if len(url) == 0 {
		return nil, errors.New("[Err] missing URL")
	}

	if len(userName) == 0 {
		return nil, errors.New("[Err] missing username")
	}

	if len(password) == 0 {
		return nil, errors.New("[Err] missing password")
	}

	var config Config
	// generate Jamf Pro Classic API Token
	credentials := []byte(userName + ":" + password)
	encodedCredentials := base64.StdEncoding.EncodeToString(credentials)
	config.classicToken = encodedCredentials

	// request Jamf Pro API Token
	req, err := http.NewRequest("POST", url + "/uapi/auth/tokens", nil)
	req.Header.Set("Authorization", "Basic " + encodedCredentials)
	if err != nil{
		fmt.Println("[Err] ", err.Error())
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[Err] ", err.Error())
	}
	defer resp.Body.Close()

	if resp.Status != "200" {
		return &config, errors.New("[Err] Request Jamf Pro API Token: HTTP Status is " + resp.Status)
	}

	var r io.Reader = resp.Body
	var v1Token V1Token
	err = json.NewDecoder(r).Decode(&v1Token)
	if err != nil {
		fmt.Println("[Err] ", err.Error())
	}
	config.v1Token = v1Token.Token

	return &config, nil
}


type Client struct {
	httpClient *http.Client
	config *Config
}


func NewClient(config *Config) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		config: config,
	}
}