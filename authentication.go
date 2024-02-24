package zoloz_go_api_sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	http_client_go "github.com/harryosmar/http-client-go"
	"time"
)

type (
	AuthenticationTestRequestBody struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
)

const (
	uriPathAuthenticationTest = "/api/v1/zoloz/authentication/test"
)

func (c zolozClient) AuthenticationTest(ctx context.Context) (*http_client_go.Response, error) {
	req := AuthenticationTestRequestBody{
		Title:       "hello",
		Description: "just for demonstration.",
	}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	headers, err := c.generateHeaders(ctx, uriPathAuthenticationTest, now, req)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(reqBytes)
	response, err := c.httpClient.Post(
		ctx,
		fmt.Sprintf("%s%s", c.baseUrl, uriPathAuthenticationTest),
		buffer,
		headers,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}
