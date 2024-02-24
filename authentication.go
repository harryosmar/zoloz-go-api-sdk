package zoloz_go_api_sdk

import (
	"context"
)

const (
	uriPathAuthenticationTest = "/api/v1/zoloz/authentication/test"
)

func (c zolozClient) AuthenticationTest(ctx context.Context) (*AuthenticationTestResponse, error) {
	return post[AuthenticationTestRequest, AuthenticationTestResponse](
		ctx,
		"zolozClient.AuthenticationTest",
		c.baseUrl,
		uriPathAuthenticationTest,
		c.clientId,
		c.signer,
		c.httpClient,
		&AuthenticationTestRequest{
			Title:       "hello",
			Description: "just for demonstration.",
		},
	)
}
