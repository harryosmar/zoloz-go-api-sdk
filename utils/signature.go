package utils

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	hash "github.com/harryosmar/hash-go"
	"net/url"
	"strings"
	"time"
)

type (
	Signature struct {
		Method      string
		Uri         string
		ClientId    string
		RequestTime time.Time
		Body        interface{}
	}
)

func LoadKeyFromBase64(base64PrivateKey string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64PrivateKey)
}

func (c Signature) String() (string, error) {
	b, err := json.Marshal(c.Body)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%s %s\n%s.%s.%s", c.Method, c.Uri, c.ClientId, GenerateZolozTimeStr(c.RequestTime), b)

	return result, nil
}

func (c Signature) Hash(ctx context.Context, signer hash.Signer) (string, error) {
	contentSignatureStr, err := c.String()
	if err != nil {
		return "", err
	}

	hashBytes, err := signer.Sign(ctx, []byte(contentSignatureStr))
	if err != nil {
		return "", err
	}

	return url.QueryEscape(
		hash.SignOutputBase64(hashBytes),
	), nil
}

func GenerateBizId() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
