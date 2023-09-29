package main

import (
	"context"
	"fmt"
	"github.com/Bostigger/go-hubtelotp/internal/hubtelotp"
	"github.com/Bostigger/go-hubtelotp/internal/util"
)

func main() {
	config := util.HubtelOTPConfig{
		ClientId:     "your_client_id",
		ClientSecret: "your_client_secret",
		ApiUrl:       "https://api-devp-otp-2704.hubtel.com/otp/send",
		Timeout:      30,
	}

	client := hubtelotp.NewHubtelOTPClient(&config)

	otpRequest := hubtelotp.OTPRequest{
		SenderId:    "your_sender_name",
		PhoneNumber: "+233xxxxxxx",
		CountryCode: "country_code",
	}

	ctx := context.Background()
	response, err := client.SendOtp(ctx, otpRequest)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(response)
}
