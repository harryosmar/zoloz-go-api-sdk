package utils_test

import (
	zoloz_go_api_sdk "github.com/harryosmar/zoloz-go-api-sdk/utils"
	"log"
	"testing"
	"time"
)

func TestGenerateZolozTimeStr(t *testing.T) {

	type args struct {
		t time.Time
	}
	testData := []struct {
		name           string
		args           args
		expectedResult string
	}{
		{
			name: "Should be valid",
			args: args{
				t: time.Unix(1554354536, 0),
			},
			expectedResult: `2019-04-04T12:08:56+0700`,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			now := time.Now()
			log.Println(now.Format(time.RFC3339))
			actual := zoloz_go_api_sdk.GenerateZolozTimeStr(tt.args.t)
			if actual != tt.expectedResult {
				t.Errorf("expect %s got %s", tt.expectedResult, actual)
				return
			}
		})
	}
}

func TestGenerateZolozTimeStrFromStr(t *testing.T) {

	type args struct {
		t string
	}
	testData := []struct {
		name           string
		args           args
		expectedResult string
	}{
		{
			name: "Given jakarta timezone",
			args: args{
				t: `2022-06-06T15:57:54+07:00`,
			},
			expectedResult: `2022-06-06T15:57:54+0700`,
		},
		{
			name: "Given UTC timezone",
			args: args{
				t: `2009-11-10T23:00:00Z`,
			},
			expectedResult: `2009-11-10T23:00:00+0000`,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			now := time.Now()
			log.Println(now.Format(time.RFC3339))
			actual := zoloz_go_api_sdk.GenerateZolozTimeStrFromStr(tt.args.t)
			if actual != tt.expectedResult {
				t.Errorf("expect %s got %s", tt.expectedResult, actual)
				return
			}
		})
	}
}
