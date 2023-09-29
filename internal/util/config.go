package util

import (
	"os"
	"strconv"
	"time"
)

type HubtelOTPConfig struct {
	ClientId     string
	ClientSecret string
	ApiUrl       string
	Timeout      time.Duration
}

func LoadConfigFromEnv() *HubtelOTPConfig {
	timeout, err := strconv.ParseInt(os.Getenv("HUBTEL_TIMEOUT"), 10, 64)
	if err != nil {
		timeout = int64(1000 * time.Second)
	}
	return &HubtelOTPConfig{
		ClientId:     os.Getenv("HUBTEL_CLIENT_ID"),
		ClientSecret: os.Getenv("HUBTEL_CLIENT_SECRET"),
		ApiUrl:       os.Getenv("HUBTEL_API_URL"),
		Timeout:      time.Duration(timeout) * time.Second,
	}
}
