package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"

	funk "github.com/thoas/go-funk"
	"net/http"
	Url "net/url"
)

const (
	AUTH_BASIC_USER_TYPE = "AUTH_BASIC_USER_TYPE"
	AUTH_BASIC_APP_TYPE  = "AUTH_BASIC_APP_TYPE"
	AUTH_BEARER_TYPE     = "AUTH_BEARER_TYPE"
)

var (
	auth   = Auth{}
	client = &http.Client{}
	// Go doesn't support constant array so using [...] to fix the size of the array
	supportedApiMethods = [...]string{"POST", "GET", "PUT"}
	supportedAuthType   = [...]string{AUTH_BASIC_APP_TYPE, AUTH_BASIC_USER_TYPE, AUTH_BEARER_TYPE}
)

type Auth struct {
	authType    string
	username    string
	password    string
	apiKey      string
	bearerToken string
}

func SetBasicUserAuth(username string, password string) {
	auth.authType = AUTH_BASIC_USER_TYPE
	auth.username = username
	auth.password = password
}

func SetBasicAppAuth(apiKey string) {
	auth.authType = AUTH_BASIC_APP_TYPE
	auth.apiKey = apiKey
}

func SetBearerAuth(bearerToken string) {
	auth.authType = AUTH_BEARER_TYPE
	auth.bearerToken = bearerToken
}

func GetAuthHeaderValue() (string, error) {
	if funk.Contains(supportedAuthType, auth.authType) == false {
		return "", errors.New("authType " + auth.authType + " is invalid.")
	}

	switch auth.authType {
	case AUTH_BASIC_APP_TYPE:
		return "Basic " + auth.apiKey, nil
	case AUTH_BASIC_USER_TYPE:
		tokenB64 := base64.StdEncoding.EncodeToString([]byte(auth.username + ":" + auth.password))
		return "Basic " + tokenB64, nil
	case AUTH_BEARER_TYPE:
		return "Bearer " + auth.bearerToken, nil
	default:
		return "", errors.New("Either authType is not set or invalid")
	}
}

type Request struct {
	body   map[string]interface{}
	header map[string]string
	method string
	url    string
}

func NewRequest() Request {
	return Request{
		body:   map[string]interface{}{},
		header: map[string]string{},
		method: "",
		url:    "",
	}
}

func (r *Request) WithBody(body map[string]interface{}) error {
	r.body = body
	return nil
}

func (r *Request) WithHeader(headerName string, headerValue string) error {
	r.header[headerName] = headerValue
	return nil
}

func (r *Request) WithMethod(method string) error {
	if funk.Contains(supportedApiMethods, method) {
		r.method = method
		return nil
	} else {
		return errors.New("Invalid method " + method)
	}
}

func (r *Request) WithUrl(url string) error {
	_, err := Url.ParseRequestURI(url)
	if err != nil {
		return err
	}
	r.url = url
	return nil
}

func (r *Request) request() (*http.Request, error) {
	var (
		err     error
		request *http.Request
	)
	if len(r.body) == 0 {
		request, err = http.NewRequest(r.method, r.url, nil)
	} else {
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(r.body)
		request, err = http.NewRequest(r.method, r.url, payloadBuf)
	}
	if err != nil {
		return &http.Request{}, err
	}

	// Set header(s) if any
	for headerName, headerValue := range r.header {
		request.Header.Add(headerName, headerValue)
	}

	return request, nil
}

func (r *Request) Send() (*http.Response, error) {
	request, err := r.request()
	if err != nil {
		return &http.Response{}, err
	}
	return client.Do(request)
}
