package utils_test

import (
	zoloz_go_api_sdk "github.com/harryosmar/zoloz-go-api-sdk/utils"
	"testing"
	"time"
)

func TestGenerateSignString(t *testing.T) {

	type args struct {
		payload zoloz_go_api_sdk.Signature
	}
	testData := []struct {
		name           string
		args           args
		expectedResult string
	}{
		{
			name: "Should be valid",
			args: args{
				payload: zoloz_go_api_sdk.Signature{
					Method:      "POST",
					Uri:         "/api/v1/zoloz/authentication/test",
					ClientId:    "9999999999999999",
					RequestTime: time.Unix(1516239022, 0),
					Body: map[string]string{
						"title":       "hello",
						"description": "just for demonstration.",
					},
				},
			},
			expectedResult: `POST /api/v1/zoloz/authentication/test
9999999999999999.2018-01-18T08:30:22+0700.{"description":"just for demonstration.","title":"hello"}`,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			signStr, err := tt.args.payload.String()
			if err != nil {
				t.Error(err)
				return
			}
			if signStr != tt.expectedResult {
				t.Errorf("expect %s got %s", tt.expectedResult, signStr)
				return
			}
		})
	}
}
