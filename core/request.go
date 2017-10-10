package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/wonderstream/twitch/logger"
)

// Request is used to manage and simplify request from API
type Request struct {
	BaseURL             string
	Method              string
	HeaderAccept        string
	HeaderClientID      string
	HeaderAuthorization string
	RequireToken        bool
	Logger              logger.Logger
}

// NewUserAccessTokenRequest constructor to build Request containing User Token
// information
func NewUserAccessTokenRequest(oauth2 *OAuth2, t TokenResponse) *Request {
	request := &Request{
		BaseURL:             oauth2.URL,
		Method:              http.MethodGet,
		HeaderAccept:        oauth2.TwitchSettings.Headers["Accept"],
		HeaderClientID:      oauth2.TwitchSettings.ClientID,
		HeaderAuthorization: t.AccessToken,
		RequireToken:        true,
	}

	return request
}

// NewRequest constructor to build a Request without any User Token information
func NewRequest(oauth2 *OAuth2) *Request {
	request := &Request{
		BaseURL:        oauth2.URL,
		Method:         http.MethodGet,
		HeaderAccept:   oauth2.TwitchSettings.Headers["Accept"],
		HeaderClientID: oauth2.TwitchSettings.ClientID,
		RequireToken:   false,
	}

	return request
}

// Compute any headers info to the request
func (r *Request) computeHeader(httpRequest *http.Request) {
	httpRequest.Header.Add("Accept", r.HeaderAccept)
	httpRequest.Header.Add("Client-ID", r.HeaderClientID)

	if r.RequireToken {
		httpRequest.Header.Add("Authorization", "OAuth "+r.HeaderAuthorization)
	}
}

// SendRequest send the request with a json structure definition be populated and returned
func (r Request) SendRequest(URI string, definition interface{}) error {
	client := &http.Client{}

	completeURL := r.BaseURL + URI
	httprRequest, _ := http.NewRequest(r.Method, completeURL, nil)
	r.computeHeader(httprRequest)

	if r.Logger != nil {
		r.Logger.LogInterface(r)
		r.Logger.LogInterface(httprRequest)
	}

	resp, err := client.Do(httprRequest)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if r.Logger != nil {
		r.Logger.LogInterface(string(body))
	}

	err = json.Unmarshal([]byte(body), &definition)
	if err != nil {
		return err
	}

	return nil
}
