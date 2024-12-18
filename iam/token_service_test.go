// © Copyright 2024 Hewlett Packard Enterprise Development LP
package iam

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"testing"

	stdlog "log"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/go-logr/stdr"
	"gotest.tools/v3/assert"
)

func Test_GetAccessToken(t *testing.T) {
	glcpUrl := "sso.common.cloud.hpe.com"
	glcpUser := "xxxxxxx-zzz-123-3456"
	glcpUserSecret := "zzzzzrandomxxxxxxzzzzz"
	proxy := "http://dummy_proxy:8080"
	log := stdr.New(stdlog.New(os.Stdout, "", stdlog.LstdFlags))

	ts := NewTokenService(glcpUrl, glcpUser, glcpUserSecret, proxy, &log)

	t.Run("GetAccessToken successful", func(t *testing.T) {
		expToken := "bearerdummyoxyzxxzzz12xxxx341111zzzzyyyyyyQQQQQHHHHHbF9fc2NGOGpT"
		oauth2 := oauth2_token{
			AccessToken: expToken,
			TokenType:   "Bearer",
			ExpiresIn:   7199,
		}
		data, _ := json.Marshal(oauth2)
		reader := bytes.NewReader(data)
		response := http.Response{
			Body: io.NopCloser(reader),
		}
		patch := gomonkey.ApplyMethodReturn(ts, "PostRequest", &response, nil)
		defer patch.Reset()

		token, err := ts.GetAccessToken()
		if err != nil {
			t.Errorf("FAILED: unexpected error")
		}
		assert.Equal(t, token, expToken)
	})

	t.Run("GetAccessToken failed", func(t *testing.T) {
		patch := gomonkey.ApplyMethodReturn(ts, "PostRequest", nil, errors.New("error while generating token"))
		defer patch.Reset()
		token, err := ts.GetAccessToken()
		if err == nil {
			t.Errorf("FAILED: expected errornot found")
		}
		assert.Equal(t, len(token), 0)
	})
}
