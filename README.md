# Go Hubtel OTP Sender

Go Hubtel OTP Sender is a small Go package that simplifies sending OTPs (One-Time Passwords) using the Hubtel OTP service. It provides an easy-to-use API for sending OTPs to phone numbers.

## Installation

To use this package in your Go project, you can simply install it using:

```
go get "github.com/Bostigger/go-hubtelotp"

Make sure you have Go modules enabled in your project.
```

## Usage
Initialize the HubtelOTPClient
```config := &util.HubtelOTPConfig{
ClientId:     "yourClientId",
ClientSecret: "yourClientSecret",
ApiUrl:       "https://api-devp-otp-2704.hubtel.com/otp/send", // or your Hubtel API URL
Timeout:      30, // Timeout in seconds
}

client := hubtelotp.NewHubtelOTPClient(config)
```

## Sending an OTP
```
otpRequest := hubtelotp.OTPRequest{
    SenderId:    "YourSenderId",
    PhoneNumber: "+233559099224",
    CountryCode: "GH",
}

response, err := client.SendOtp(context.Background(), otpRequest)
if err != nil {
    // Handle error
    fmt.Println("Error:", err)
    return
}

fmt.Println("OTP Sent. Response:", response)

```
## Configuration
You can configure the HubtelOTPClient either by providing the configuration directly or by using environment variables. Here are the available configuration options:

```
HUBTEL_CLIENT_ID: Your Hubtel Client ID.
HUBTEL_CLIENT_SECRET: Your Hubtel Client Secret.
HUBTEL_API_URL: The Hubtel OTP API URL.
HUBTEL_TIMEOUT: Timeout for API requests in seconds.
```

### Contributing
If you'd like to contribute to this project, please follow these steps:

Fork the project.
Create a new branch for your feature or bug fix.
Make your changes and add tests if applicable.
Ensure all tests pass.
Submit a pull request.