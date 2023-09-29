package hubtelotp_test

import (
	"context"
	"github.com/Bostigger/go-hubtelotp/internal/hubtelotp" // replace with your package path
	"github.com/Bostigger/go-hubtelotp/internal/util"      // replace with your package path
	"testing"
)

func TestNewHubtelOTPClient(t *testing.T) {
	config := &util.HubtelOTPConfig{
		ClientId:     "testClientId",
		ClientSecret: "testClientSecret",
		ApiUrl:       "testUrl",
	}

	client := hubtelotp.NewHubtelOTPClient(config)
	if client == nil {
		t.Fatal("Expected HubtelOTPClient to be created, got nil")
	}
}

func TestSendOtp(t *testing.T) {
	config := &util.HubtelOTPConfig{
		ClientId:     "testClientId",
		ClientSecret: "testClientSecret",
		ApiUrl:       "testUrl",
	}

	client := hubtelotp.NewHubtelOTPClient(config)

	request := hubtelotp.OTPRequest{
		SenderId:    "test",
		PhoneNumber: "testNumber",
		CountryCode: "GH",
	}

	ctx := context.Background()

	//TODO :  mock the actual API call here

	response, err := client.SendOtp(ctx, request)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// 	//TODO :  add more expected res
	if response.Code != "expectedCode" {
		t.Fatalf("Expected response code %s, got %s", "expectedCode", response.Code)
	}
}
