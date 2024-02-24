package zoloz_go_api_sdk

import (
	"context"
	"fmt"
	hash "github.com/harryosmar/hash-go"
	http_client "github.com/harryosmar/http-client-go"
	"github.com/harryosmar/zoloz-go-api-sdk/utils"
	"net/http"
	"time"
)

type ZolozClient interface {
	AuthenticationTest(ctx context.Context) (*http_client.Response, error)
	RealIdInit(ctx context.Context, req *RealIdInitRequest) (*RealIdInitResponse, error)
	RealIdCheckResult(ctx context.Context, req *RealIdCheckResultRequest) (*RealIdCheckResultResponse, error)
}

type (
	zolozClient struct {
		baseUrl    string
		httpClient http_client.HttpClientRepository
		clientId   string
		signer     hash.Signer
	}
)

func NewZolozClient(baseUrl string, httpClient http_client.HttpClientRepository, clientId string, signer hash.Signer) *zolozClient {
	return &zolozClient{baseUrl: baseUrl, httpClient: httpClient, clientId: clientId, signer: signer}
}

func (c zolozClient) generateHeaders(ctx context.Context, uri string, now time.Time, req interface{}) (map[string]string, error) {
	signature, err := utils.Signature{
		Method:      http.MethodPost,
		Uri:         uri,
		ClientId:    c.clientId,
		RequestTime: now,
		Body:        req,
	}.Hash(ctx, c.signer)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"Content-Type": "application/json; charset=UTF-8",
		"Client-Id":    c.clientId,
		"Request-Time": utils.GenerateZolozTimeStr(now),
		"Signature":    fmt.Sprintf("algorithm=RSA256, signature=%s", signature),
	}, nil
}
