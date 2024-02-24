package zoloz_go_api_sdk

import (
	"context"
)

const (
	uriPathRealIdInit        = "/api/v1/zoloz/realid/initialize"
	uriPathRealIdCheckResult = "/api/v1/zoloz/realid/checkresult"
)

func (c zolozClient) RealIdInit(ctx context.Context, reqRaw *RealIdInitRequest) (*RealIdInitResponse, error) {
	return post[RealIdInitRequest, realIdInitRequest, RealIdInitResponse](
		ctx,
		"zolozClient.RealIdInit",
		c.baseUrl,
		uriPathRealIdInit,
		c.clientId,
		c.signer,
		c.httpClient,
		reqRaw,
	)
}

func (c zolozClient) RealIdCheckResult(ctx context.Context, reqRaw *RealIdCheckResultRequest) (*RealIdCheckResultResponse, error) {
	return post[RealIdCheckResultRequest, realIdCheckResultRequest, RealIdCheckResultResponse](
		ctx,
		"zolozClient.RealIdCheckResult",
		c.baseUrl,
		uriPathRealIdCheckResult,
		c.clientId,
		c.signer,
		c.httpClient,
		reqRaw,
	)
}
