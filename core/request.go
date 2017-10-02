package core

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Request is used to manage and simplify request from API
type Request struct {
	BaseURL             string
	Method              string
	HeaderAccept        string
	HeaderClientID      string
	HeaderAuthorization string
}

// NewRequest constructor
func NewRequest(oauth2 *OAuth2, t TokenResponse) *Request {
	request := &Request{
		BaseURL:             oauth2.URL,
		HeaderAccept:        oauth2.TwitchSettings.Headers["Accept"],
		HeaderClientID:      oauth2.TwitchSettings.ClientID,
		HeaderAuthorization: t.AccessToken,
	}

	return request
}

// Compute any headers info to the request
func (r *Request) computeHeader(httpRequest *http.Request) {
	httpRequest.Header.Add("Accept", r.HeaderAccept)
	httpRequest.Header.Add("Client-ID", r.HeaderClientID)
	httpRequest.Header.Add("Authorization", "OAuth "+r.HeaderAuthorization)
}

// Request send the request with a json structure definition be populated and returned
func (r Request) sendRequest(URI string, definition interface{}) error {
	client := &http.Client{}

	completeURL := r.BaseURL + URI
	httprRequest, _ := http.NewRequest(r.Method, completeURL, nil)
	r.computeHeader(httprRequest)

	resp, err := client.Do(httprRequest)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(body), &definition)
	if err != nil {
		return err
	}

	return nil
}
