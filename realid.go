package zoloz_go_api_sdk

import (
	"context"
	"github.com/harryosmar/zoloz-go-api-sdk/enums"
)

type (
	PageConfig struct {
		UrlFaceGuide string `json:"urlFaceGuide"`
	}

	ProductConfig struct {
		CropDocImage       enums.YesNo
		CropFaceImage      enums.YesNo
		LandmarkCheck      enums.LandmarkChecks
		HologramCheck      enums.HologramChecks
		PageInfoCheck      enums.PageInfoChecks
		PreciseTamperCheck enums.YesNo
		DocUiType          enums.DocUiType
		SpoofMode          enums.Level
		LivenessMode       enums.Level
		AntiInjectionMode  enums.Level
		ActionCheckItems   enums.ActionsCheckItem
		ActionRandom       enums.YesNo
		ActionFrame        enums.ActionsFrame
		RiskMode           enums.Level
		IdnThreshold       int32
	}

	H5ModeConfig struct {
		CompleteCallbackUrl  string `json:"completeCallbackUrl"`
		InterruptCallbackUrl string `json:"interruptCallbackUrl"`
		State                string `json:"state"`
		IsIframe             string `json:"isIframe"`
		Locale               string `json:"locale"`
	}

	productConfig struct {
		CropDocImage       string         `json:"cropDocImage,omitempty"`
		CropFaceImage      string         `json:"cropFaceImage,omitempty"`
		LandmarkCheck      map[string]any `json:"landmarkCheck,omitempty"`
		HologramCheck      map[string]any `json:"hologramCheck,omitempty"`
		PageInfoCheck      map[string]any `json:"pageInfoCheck,omitempty"`
		PreciseTamperCheck string         `json:"preciseTamperCheck,omitempty"`
		DocUiType          string         `json:"docUiType,omitempty"`
		SpoofMode          string         `json:"spoofMode,omitempty"`
		LivenessMode       string         `json:"livenessMode,omitempty"`
		AntiInjectionMode  string         `json:"antiInjectionMode,omitempty"`
		ActionCheckItems   []string       `json:"actionCheckItems,omitempty"`
		ActionRandom       string         `json:"actionRandom,omitempty"`
		ActionFrame        []string       `json:"actionFrame,omitempty"`
		RiskMode           string         `json:"riskMode,omitempty"`
		IdnThreshold       int32          `json:"idnThreshold,omitempty"`
	}

	realIdInitRequest struct {
		BizId                string         `json:"bizId"`
		FlowType             string         `json:"flowType"`
		UserId               string         `json:"userId"`
		DocType              string         `json:"docType"`
		PageConfig           *PageConfig    `json:"pageConfig,omitempty"`
		ServiceLevel         string         `json:"serviceLevel,omitempty"`
		OperationMode        string         `json:"operationMode,omitempty"`
		MetaInfo             string         `json:"metaInfo"`
		SceneCode            string         `json:"sceneCode,omitempty"`
		Pages                string         `json:"pages,omitempty"`
		H5ModeConfig         *H5ModeConfig  `json:"h5ModeConfig,omitempty"`
		ProductConfig        *productConfig `json:"productConfig,omitempty"`
		AllowExpiredDocument string         `json:"allowExpiredDocument,omitempty"`
	}

	RealIdInitRequest struct {
		BizId                string
		FlowType             enums.FlowType
		UserId               string
		DocType              enums.DocType
		PageConfig           *PageConfig
		ServiceLevel         enums.ServiceLevel
		OperationMode        enums.OperationMode
		MetaInfo             string
		SceneCode            string
		Pages                string
		H5ModeConfig         *H5ModeConfig
		ProductConfig        *ProductConfig
		AllowExpiredDocument enums.YesNo
	}

	Result struct {
		ResultStatus  string `json:"resultStatus"`
		ResultCode    string `json:"resultCode"`
		ResultMessage string `json:"resultMessage"`
	}

	RealIdInitResponse struct {
		Result        Result `json:"result"`
		TransactionId string `json:"transactionId"`
		ClientCfg     string `json:"clientCfg"`
	}

	realIdCheckResultRequest struct {
		BizId                         string   `json:"bizId"`
		TransactionId                 string   `json:"transactionId"`
		IsReturnImage                 string   `json:"isReturnImage,omitempty"`
		ExtraImageControlList         []string `json:"extraImageControlList,omitempty"`
		ReturnFiveCategorySpoofResult string   `json:"returnFiveCategorySpoofResult,omitempty"`
	}

	RealIdCheckResultRequest struct {
		BizId                         string
		TransactionId                 string
		IsReturnImage                 enums.YesNo
		ExtraImageControlList         enums.ExtraImageControlList
		ReturnFiveCategorySpoofResult enums.YesNo
	}

	ExtBasicInfo struct {
		CertType string `json:"certType"`
		CertNo   string `json:"certNo"`
		CertName string `json:"certName"`
	}

	ExtFaceInfo struct {
		EkycResultFace     string            `json:"ekycResultFace"`
		FaceScore          int               `json:"faceScore"`
		FaceImg            string            `json:"faceImg"`
		ExtraImages        map[string]string `json:"extraImages"`
		FaceQuality        float32           `json:"faceQuality"`
		FaceLivenessResult string            `json:"faceLivenessResult"`
		EstimatedAge       int               `json:"estimatedAge"`
	}

	ExtraSpoofResultDetailComponents struct {
		Name      string `json:"name"`
		SubResult string `json:"subResult"`
	}

	ExtraSpoofResultDetail struct {
		Name       string                             `json:"name"`
		Result     string                             `json:"result"`
		SpoofType  string                             `json:"spoofType"`
		Components []ExtraSpoofResultDetailComponents `json:"components"`
	}

	ExtIdInfo struct {
		EkycResultDoc           string            `json:"ekycResultDoc"`
		DocEdition              int               `json:"docEdition"`
		FrontPageImg            string            `json:"frontPageImg"`
		BackPageImg             string            `json:"backPageImg"`
		ExtraImages             map[string]string `json:"extraImages"`
		OcrResult               map[string]interface{}
		SpoofResult             map[string]interface{}
		ExtraSpoofResultDetails []ExtraSpoofResultDetail `json:"extraSpoofResultDetails"`
		SecurityFeaturesResult  map[string]interface{}   `json:"securityFeaturesResult"`
		DocErrorDetails         string                   `json:"docErrorDetails"`
	}

	ExtRiskInfo struct {
		EkycResultRisk         string `json:"ekycResultRisk"`
		StrategyPassResult     string `json:"strategyPassResult"`
		IdNetworkDetails       string `json:"idNetworkDetails"`
		OtherRiskReasonDetails string `json:"otherRiskReasonDetails"`
	}

	IdNetworkDetails struct {
		EkycId     string   `json:"ekyc_id"`
		Mobile     string   `json:"mobile"`
		ReasonCode []string `json:"reason_code"`
		Type       string   `json:"type"`
		Score      string   `json:"score"`
		UserId     string   `json:"user_id"`
	}

	RealIdCheckResultResponse struct {
		Result       Result       `json:"result"`
		EkycResult   string       `json:"ekycResult"`
		ExtBasicInfo ExtBasicInfo `json:"extBasicInfo"`
		ExtFaceInfo  ExtFaceInfo  `json:"extFaceInfo"`
		ExtIdInfo    ExtIdInfo    `json:"extIdInfo"`
		ExtRiskInfo  ExtRiskInfo  `json:"extRiskInfo"`
	}

	Request[In any, Out any] interface {
		Convert() *Out
	}
)

func (p *ProductConfig) Convert() *productConfig {
	if p == nil {
		return nil
	}

	return &productConfig{
		CropDocImage:       p.CropDocImage.String(),
		CropFaceImage:      p.CropFaceImage.String(),
		LandmarkCheck:      p.LandmarkCheck.ToMap(),
		HologramCheck:      p.HologramCheck.ToMap(),
		PageInfoCheck:      p.PageInfoCheck.ToMap(),
		PreciseTamperCheck: p.PreciseTamperCheck.String(),
		DocUiType:          p.DocUiType.String(),
		SpoofMode:          p.SpoofMode.String(),
		LivenessMode:       p.LivenessMode.String(),
		AntiInjectionMode:  p.AntiInjectionMode.String(),
		ActionCheckItems:   p.ActionCheckItems.ToStrings(),
		ActionRandom:       p.ActionRandom.String(),
		ActionFrame:        p.ActionFrame.ToStrings(),
		RiskMode:           p.RiskMode.String(),
		IdnThreshold:       p.IdnThreshold,
	}
}

func (req *RealIdInitRequest) Convert() *realIdInitRequest {
	if req == nil {
		return &realIdInitRequest{}
	}

	return &realIdInitRequest{
		BizId:                req.BizId,
		MetaInfo:             req.MetaInfo,
		FlowType:             req.FlowType.String(),
		DocType:              req.DocType.String(),
		UserId:               req.UserId,
		SceneCode:            req.SceneCode,
		ServiceLevel:         req.ServiceLevel.String(),
		OperationMode:        req.OperationMode.String(),
		Pages:                req.Pages,
		PageConfig:           req.PageConfig,
		H5ModeConfig:         req.H5ModeConfig,
		ProductConfig:        req.ProductConfig.Convert(),
		AllowExpiredDocument: req.AllowExpiredDocument.String(),
	}
}

func (req *RealIdCheckResultRequest) Convert() *realIdCheckResultRequest {
	if req == nil {
		return &realIdCheckResultRequest{}
	}

	return &realIdCheckResultRequest{
		BizId:                         req.BizId,
		TransactionId:                 req.TransactionId,
		IsReturnImage:                 req.IsReturnImage.String(),
		ExtraImageControlList:         req.ExtraImageControlList.ToStringList(),
		ReturnFiveCategorySpoofResult: req.ReturnFiveCategorySpoofResult.String(),
	}
}

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
