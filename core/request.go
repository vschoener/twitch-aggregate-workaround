package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/wonderstream/twitch/logger"
)

// Request is used to manage and simplify request from API
type Request struct {
	BaseURL string
	Method  string
	Headers map[string]string
	Form    map[string]string
	Logger  logger.Logger
}

// NewRequest constructor to build a Request without any User Token information
func NewRequest(oauth2 *OAuth2, t *TokenResponse) *Request {
	request := &Request{
		BaseURL: oauth2.URL,
		Method:  http.MethodGet,
		Headers: make(map[string]string),
		Form:    make(map[string]string),
	}

	request.Headers["Accept"] = oauth2.TwitchSettings.Headers["Accept"]
	request.Headers["Client-ID"] = oauth2.TwitchSettings.ClientID

	if nil != t && t.IsAuthenticated() {
		request.Headers["Authorization"] = "OAuth " + t.AccessToken
	}

	return request
}

// Compute any headers info to the request
func (r *Request) computeHeader(httpRequest *http.Request) {
	for name, value := range r.Headers {
		httpRequest.Header.Set(name, value)
	}
}

// SetPost prepare Post data request
func (r *Request) SetPost(data map[string]string, contentType string) {
	r.Headers["Content-Type"] = contentType
	r.Method = http.MethodPost
	r.Form = data
}

// SendRequest send the request with a json structure definition be populated and returned
func (r Request) SendRequest(URI string, definition interface{}) error {
	client := &http.Client{}

	jsonRaw := make([]byte, 0)
	if r.Method == http.MethodPost {
		jsonRaw, _ = json.Marshal(r.Form)
	}

	completeURL := r.BaseURL + URI
	request, _ := http.NewRequest(r.Method, completeURL, bytes.NewBuffer(jsonRaw))
	r.computeHeader(request)

	r.Logger.Log(fmt.Sprintf("Request on %s, %#v", completeURL, r))

	resp, err := client.Do(request)
	if err != nil {
		r.Logger.Log(fmt.Sprintf("Error %s", err))
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.Logger.Log(fmt.Sprintf("Error %s, response HTTP %d with Body %s", err, resp.StatusCode, string(body)))
		return err
	} else if resp.StatusCode != http.StatusOK {
		r.Logger.Log(fmt.Sprintf("Error status code = %d, response HTTP %d with Body %s", resp.StatusCode, resp.StatusCode, string(body)))
		return errors.New("Status code errror")
	}

	r.Logger.Log(fmt.Sprintf("Response HTTP %d with Body %s", resp.StatusCode, string(body)))

	err = json.Unmarshal([]byte(body), &definition)
	if err != nil {
		r.Logger.Log(fmt.Sprintf("Json unmarshal error %s", err))
		return err
	}

	return nil
}
