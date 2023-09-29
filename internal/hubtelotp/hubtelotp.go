package hubtelotp

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/Bostigger/go-hubtelotp/internal/util"
	"net/http"
)

type OTPRequest struct {
	SenderId    string `json:"senderId"`
	PhoneNumber string `json:"phoneNumber"`
	CountryCode string `json:"countryCode"`
}

type OTPResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Data    struct {
		RequestId string `json:"requestId"`
		Prefix    string `json:"prefix"`
	}
}

type HubtelOTPClient struct {
	Config util.HubtelOTPConfig
}

func NewHubtelOTPClient(config *util.HubtelOTPConfig) *HubtelOTPClient {
	if config == nil {
		config = util.LoadConfigFromEnv()
	}
	return &HubtelOTPClient{
		Config: *config,
	}
}

func (client *HubtelOTPClient) SendOtp(ctx context.Context, otpRequest OTPRequest) (OTPResponse, error) {
	requestBody, err := json.Marshal(otpRequest)
	if err != nil {
		return OTPResponse{}, err
	}
	//create http request
	req, err := http.NewRequestWithContext(ctx, "POST", client.Config.ApiUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return OTPResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(client.Config.ClientId+":"+client.Config.ClientSecret)))

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return OTPResponse{}, err
	}
	defer resp.Body.Close()

	//Decode the response

	var otpResponse OTPResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&otpResponse); err != nil {
		return OTPResponse{}, err
	}

	return otpResponse, nil
}
