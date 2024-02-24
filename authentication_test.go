package zoloz_go_api_sdk_test

import (
	"context"
	hash "github.com/harryosmar/hash-go"
	http_client_go "github.com/harryosmar/http-client-go"
	zoloz_go_api_sdk "github.com/harryosmar/zoloz-go-api-sdk"
	"net/http"
	"os"
	"testing"
	"time"
)

func getClient() zoloz_go_api_sdk.ZolozClient {
	privateKey, err := hash.LoadPrivateKeyFromBase64Encoded(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	return zoloz_go_api_sdk.NewZolozClient(
		os.Getenv("ZOLOZ_BASE_URL_ID"),
		http_client_go.NewHttpClientRepository(
			&http.Client{
				Timeout: 5 * time.Second,
			},
		).EnableDebug(),
		os.Getenv("ZOLOZ_CLIENT_ID"),
		hash.NewRsa256Signer(privateKey),
	)
}

func TestAuth(t *testing.T) {

	type args struct {
	}
	testData := []struct {
		name           string
		args           args
		expectedResult string
	}{
		{
			name: "Should be valid",
			args: args{},
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			client := getClient()

			_, err := client.AuthenticationTest(context.TODO())
			if err != nil {
				t.Errorf("AuthenticationTest err %v", err)
				return
			}
		})
	}
}
