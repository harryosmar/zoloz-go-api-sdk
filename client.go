package zoloz_go_api_sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	hash "github.com/harryosmar/hash-go"
	http_client "github.com/harryosmar/http-client-go"
	"github.com/harryosmar/zoloz-go-api-sdk/utils"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type ZolozClient interface {
	AuthenticationTest(ctx context.Context) (*AuthenticationTestResponse, error)
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
	logrus.SetFormatter(&logrus.JSONFormatter{})
	return &zolozClient{baseUrl: baseUrl, httpClient: httpClient, clientId: clientId, signer: signer}
}

func post[
	ReqT AuthenticationTestRequest | realIdInitRequest | realIdCheckResultRequest,
	ResT AuthenticationTestResponse | RealIdInitResponse | RealIdCheckResultResponse,
](
	ctx context.Context,
	fnName string,
	baseUrl string,
	urlPath string,
	clientId string,
	signer hash.Signer,
	httpClient http_client.HttpClientRepository,
	req *ReqT,
) (*ResT, error) {
	if ctx.Value(http_client.XRequestIdContext) == nil {
		ctx = context.WithValue(context.TODO(), http_client.XRequestIdContext, uuid.New().String())
	}

	var (
		err error
	)

	defer func() {
		if err != nil {
			logrus.Errorf("%s got err %v", fnName, err)
		}
	}()

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	headers, err := utils.Signature{
		Method:      http.MethodPost,
		Uri:         urlPath,
		ClientId:    clientId,
		RequestTime: now,
		Body:        req,
	}.GenerateHeaders(ctx, signer, now)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(reqBytes)
	response, err := httpClient.Post(
		ctx,
		fmt.Sprintf("%s%s", baseUrl, urlPath),
		buffer,
		headers,
	)
	if err != nil {
		return nil, err
	}

	if response.Status != 200 {
		err = fmt.Errorf("response Status Code != 200, %+v", response)
		return nil, err
	}

	var resp ResT
	err = json.Unmarshal(response.Content, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c zolozClient) generateHeaders(ctx context.Context, uri string, now time.Time, req interface{}) (map[string]string, error) {
	return utils.Signature{
		Method:      http.MethodPost,
		Uri:         uri,
		ClientId:    c.clientId,
		RequestTime: now,
		Body:        req,
	}.GenerateHeaders(ctx, c.signer, now)
}
