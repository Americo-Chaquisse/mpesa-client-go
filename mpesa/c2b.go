package mpesa

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
)

type C2BRequest struct {
	TransactionReference string
	ThirdPartyReference  string
	CustomerMSISDN       string
	Amount               string
}

type C2BResponse struct {
	Success             bool
	ThirdPartyReference string
	ConversationId      string
	TransactionId       string
	ResponseCode        string
	ResponseDescription string
}

var restClient = resty.New()

// C2B Customer to Business
func (client *Client) C2B(params C2BRequest) (C2BResponse, error) {
	client.Config.SetDefaults()
	// Get the auth token
	token, err := AuthToken(client.Config.PublicKey, client.Config.ApiKey)
	if err != nil {
		return C2BResponse{}, err
	}

	// compose request body
	body := map[string]interface{}{
		"input_TransactionReference": params.TransactionReference,
		"input_CustomerMSISDN":       params.CustomerMSISDN,
		"input_Amount":               params.Amount,
		"input_ThirdPartyReference":  params.ThirdPartyReference,
		"input_ServiceProviderCode":  client.Config.ServiceProviderCode,
	}

	// Make the request
	resp, err := restClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Origin", client.Config.origin).
		SetAuthToken(token).
		SetBody(body).
		Post(fmt.Sprintf("%s/ipg/v1x/c2bPayment/singleStage/", client.Config.Host))

	log.Printf("MPESA_REQUEST_DURATION: %v seconds", resp.Time().Seconds())

	if err != nil {
		log.Printf("MPESA_ERROR: %v, HTTP Response: %v", err, resp.String())
		return C2BResponse{}, err
	}

	// Parse request body
	var response = make(map[string]string)
	err1 := json.Unmarshal([]byte(resp.String()), &response)
	if err1 != nil {
		log.Printf("MPESA_PARSE_ERROR: %v, HTTP Response: %v", err1, resp.String())
		return C2BResponse{}, err1
	}

	log.Printf("MPESA_RESPONSE: %v", resp.String())

	paymentResponse := C2BResponse{
		ThirdPartyReference: response["output_ThirdPartyReference"],
		ConversationId:      response["output_ConversationID"],
		TransactionId:       response["output_TransactionID"],
		ResponseCode:        response["output_ResponseCode"],
		ResponseDescription: response["output_ResponseDesc"],
	}

	paymentResponse.Success = paymentResponse.ResponseCode == "INS-0"

	return paymentResponse, nil
}
