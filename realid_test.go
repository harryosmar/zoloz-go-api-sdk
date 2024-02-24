package zoloz_go_api_sdk_test

import (
	"context"
	"encoding/json"
	zoloz_go_api_sdk "github.com/harryosmar/zoloz-go-api-sdk"
	"github.com/harryosmar/zoloz-go-api-sdk/enums"
	"github.com/harryosmar/zoloz-go-api-sdk/utils"
	"testing"
	"time"
)

func TestRealId(t *testing.T) {

	type args struct {
		initRequest *zoloz_go_api_sdk.RealIdInitRequest
	}
	testData := []struct {
		name           string
		args           args
		expectedResult string
	}{
		{
			name: "Should be valid",
			args: args{
				initRequest: &zoloz_go_api_sdk.RealIdInitRequest{
					BizId: utils.GenerateBizId(),
					MetaInfo: func() string {
						meta := map[string]interface{}{
							"deviceType":  "android",
							"appVersion":  "1.0",
							"appName":     "makcik",
							"osVersion":   "7.1.1",
							"bioMetaInfo": "3.37.0:262144,0",
							"apdidToken":  "mock-apdidToken",
							"deviceModel": "MI 6",
							"zimVer":      "2.0.0",
						}
						bytes, _ := json.Marshal(meta)
						return string(bytes)
					}(),
					FlowType:      enums.FlowType_REALIDLITE_KYC,
					DocType:       enums.DocType_IndonesiaEKtp,
					UserId:        "123456abcd",
					ServiceLevel:  enums.ServiceLevel_REALID0001,
					OperationMode: enums.OperationMode_STANDARD,
					PageConfig: &zoloz_go_api_sdk.PageConfig{
						UrlFaceGuide: "http://url-to-face-guide.htm",
					},
				},
			},
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			client := getClient()

			ctx := context.TODO()
			initResponse, err := client.RealIdInit(ctx, tt.args.initRequest)
			if err != nil {
				t.Error(err)
				return
			}
			time.Sleep(1 * time.Second)

			_, err = client.RealIdCheckResult(ctx, &zoloz_go_api_sdk.RealIdCheckResultRequest{
				BizId:         tt.args.initRequest.BizId,
				TransactionId: initResponse.TransactionId,
				//TransactionId: "G000000006FRL20240224000000192190740060",
				IsReturnImage: enums.YesNo_Y,
				ExtraImageControlList: []enums.ExtraImageControl{
					enums.ExtraImageControl_FACE_EYE_CLOSE,
					enums.ExtraImageControl_DOC_FRONT_ANGLE,
				},
				ReturnFiveCategorySpoofResult: enums.YesNo_Y,
			})
			if err != nil {
				t.Error(err)
				return
			}
		})
	}
}
