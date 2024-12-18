// © Copyright 2024 Hewlett Packard Enterprise Development LP
package iam

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-logr/logr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HPE GreenLake API Gateway Service Provisioner

var errEmptyToken = errors.New("empty token received from Green Lake Cloud Platform")

// Defines structure for the oauth2_token response provided by GLCP
type oauth2_token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// Defines attributes to create an GLCP Token instance
// These attributes are used to provide GLCP credentials to the DSCC
type token_service struct {
	glcpCloudUrl    string
	apiCLientId     string
	apiClientSecret string
	proxy           string
	log             *logr.Logger
}

// Creates an instance of the TokenService struct
func NewTokenService(glcpUrl, clientId, secret, proxy string, log *logr.Logger) *token_service {
	return &token_service{
		glcpCloudUrl:    glcpUrl,
		apiCLientId:     clientId,
		apiClientSecret: secret,
		proxy:           proxy,
		log:             log,
	}
}

func (t *token_service) string() (s string) {
	return fmt.Sprintf("grant_type=%s&client_id=%s&client_secret=%s", ACCESS_GRANT_TYPE, t.apiCLientId, t.apiClientSecret)
}

// Creates a fresh bearer token for the GLCP API user
func (t *token_service) GetAccessToken() (string, error) {
	h := map[string]string{"Content-Type": APPLICATION_URL_ENCODED}
	cloudUrl := fmt.Sprintf("%s/%s", t.glcpCloudUrl, GLCP_ACCESS_TOKEN_URL)
	if !strings.Contains(cloudUrl, "https") {
		cloudUrl = "https://" + cloudUrl
	}
	res, err := t.PostRequest(cloudUrl, bytes.NewReader([]byte(t.string())), h)

	if err != nil {
		t.log.Error(err, "Error when calling `Token REST API`. Please verify the connectivity to DSCC.")
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.log.Error(err, "Unable to retrieve token from `Token REST API response`")
		return "", err
	}

	var o1 oauth2_token

	if err = json.Unmarshal(body, &o1); err != nil {
		t.log.Error(err, "Failed to unmarshal token from `Token REST API response`")
		return "", err
	} else if len(o1.AccessToken) == 0 {
		t.log.Error(errEmptyToken, "error while retrieving GLCP API client token")
		return "", status.Error(codes.InvalidArgument, "failed to generate token from Green Lake Cloud Platform. Please verify the correctness of GLCP information passed within the secret.")
	}

	return o1.AccessToken, err

}

// Visible & to be directly used only for UT
func (t *token_service) PostRequest(uri string, body *bytes.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, uri, body)
	if err != nil {
		t.log.Error(err, "client: could not create request")
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := http.Client{
		Timeout: HTTP_REQUEST_TIMEOUT,
	}
	if len(t.proxy) != 0 {
		proxy, err := url.Parse(t.proxy)
		if err != nil {
			t.log.Error(err, "error parsing proxy: "+t.proxy)
			return nil, err
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	}
	return client.Do(req)
}
