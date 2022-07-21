package mpesa

import (
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestC2BSuccess(t *testing.T) {
	httpmock.ActivateNonDefault(restClient.GetClient())
	defer httpmock.DeactivateAndReset()

	// Mock C2B Endpoint
	httpmock.RegisterResponder("POST", "https://api.sandbox.vm.co.mz:18352/ipg/v1x/c2bPayment/singleStage/",
		httpmock.NewStringResponder(200, `{"output_ResponseCode":"INS-0","output_ResponseDesc":"Request processed successfully","output_TransactionID":"8loxwg1xwt4c","output_ConversationID":"dfc12eec891244c6847b4ea594496bc3","output_ThirdPartyReference":"3QYWDW"}`))

	config := Config{
		PublicKey:           "MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAmptSWqV7cGUUJJhUBxsMLonux24u+FoTlrb+4Kgc6092JIszmI1QUoMohaDDXSVueXx6IXwYGsjjWY32HGXj1iQhkALXfObJ4DqXn5h6E8y5/xQYNAyd5bpN5Z8r892B6toGzZQVB7qtebH4apDjmvTi5FGZVjVYxalyyQkj4uQbbRQjgCkubSi45Xl4CGtLqZztsKssWz3mcKncgTnq3DHGYYEYiKq0xIj100LGbnvNz20Sgqmw/cH+Bua4GJsWYLEqf/h/yiMgiBbxFxsnwZl0im5vXDlwKPw+QnO2fscDhxZFAwV06bgG0oEoWm9FnjMsfvwm0rUNYFlZ+TOtCEhmhtFp+Tsx9jPCuOd5h2emGdSKD8A6jtwhNa7oQ8RtLEEqwAn44orENa1ibOkxMiiiFpmmJkwgZPOG/zMCjXIrrhDWTDUOZaPx/lEQoInJoE2i43VN/HTGCCw8dKQAwg0jsEXau5ixD0GUothqvuX3B9taoeoFAIvUPEq35YulprMM7ThdKodSHvhnwKG82dCsodRwY428kg2xM/UjiTENog4B6zzZfPhMxFlOSFX4MnrqkAS+8Jamhy1GgoHkEMrsT5+/ofjCx0HjKbT5NuA2V/lmzgJLl3jIERadLzuTYnKGWxVJcGLkWXlEPYLbiaKzbJb2sYxt+Kt5OxQqC1MCAwEAAQ==",
		ApiKey:              "enkrsco5guq3g57ypzc1g0yt2djexi50",
		ServiceProviderCode: "171717",
	}
	client := Client{Config: config}

	request := C2BRequest{
		TransactionReference: "DONOR12345",
		ThirdPartyReference:  "3QYWDW",
		CustomerMSISDN:       "258842058817",
		Amount:               "10",
	}

	response, err := client.C2B(request)

	if err != nil {
		t.Fatal(err)
	}

	if response.ThirdPartyReference != request.ThirdPartyReference {
		t.Fatalf("Mimatch thirdy party reference. Expected: %s, Got: %s",
			request.ThirdPartyReference, response.ThirdPartyReference)
	}

	if response.ResponseCode != "INS-0" {
		t.Fatalf("Transaction not processed successfully. Code: %s, Description: %s",
			response.ResponseCode, response.ResponseDescription)
	}
}

func TestC2BDuplicateTransaction(t *testing.T) {
	httpmock.ActivateNonDefault(restClient.GetClient())
	defer httpmock.DeactivateAndReset()

	// Mock C2B Endpoint
	httpmock.RegisterResponder("POST", "https://api.sandbox.vm.co.mz:18352/ipg/v1x/c2bPayment/singleStage/",
		httpmock.NewStringResponder(200, `{"output_ResponseCode":"INS-10","output_ResponseDesc":"Duplicate Transaction","output_TransactionID":"8loxwg1xwt4c","output_ConversationID":"dfc12eec891244c6847b4ea594496bc3","output_ThirdPartyReference":"3QYWDW"}`))

	config := Config{
		PublicKey:           "MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAmptSWqV7cGUUJJhUBxsMLonux24u+FoTlrb+4Kgc6092JIszmI1QUoMohaDDXSVueXx6IXwYGsjjWY32HGXj1iQhkALXfObJ4DqXn5h6E8y5/xQYNAyd5bpN5Z8r892B6toGzZQVB7qtebH4apDjmvTi5FGZVjVYxalyyQkj4uQbbRQjgCkubSi45Xl4CGtLqZztsKssWz3mcKncgTnq3DHGYYEYiKq0xIj100LGbnvNz20Sgqmw/cH+Bua4GJsWYLEqf/h/yiMgiBbxFxsnwZl0im5vXDlwKPw+QnO2fscDhxZFAwV06bgG0oEoWm9FnjMsfvwm0rUNYFlZ+TOtCEhmhtFp+Tsx9jPCuOd5h2emGdSKD8A6jtwhNa7oQ8RtLEEqwAn44orENa1ibOkxMiiiFpmmJkwgZPOG/zMCjXIrrhDWTDUOZaPx/lEQoInJoE2i43VN/HTGCCw8dKQAwg0jsEXau5ixD0GUothqvuX3B9taoeoFAIvUPEq35YulprMM7ThdKodSHvhnwKG82dCsodRwY428kg2xM/UjiTENog4B6zzZfPhMxFlOSFX4MnrqkAS+8Jamhy1GgoHkEMrsT5+/ofjCx0HjKbT5NuA2V/lmzgJLl3jIERadLzuTYnKGWxVJcGLkWXlEPYLbiaKzbJb2sYxt+Kt5OxQqC1MCAwEAAQ==",
		ApiKey:              "enkrsco5guq3g57ypzc1g0yt2djexi50",
		ServiceProviderCode: "171717",
	}
	client := Client{Config: config}

	request := C2BRequest{
		TransactionReference: "DONOR12345",
		ThirdPartyReference:  "3QYWDW",
		CustomerMSISDN:       "258842058817",
		Amount:               "10",
	}

	response, err := client.C2B(request)

	if err != nil {
		t.Fatal(err)
	}

	if response.ThirdPartyReference != request.ThirdPartyReference {
		t.Fatalf("Mimatch thirdy party reference. Expected: %s, Got: %s",
			request.ThirdPartyReference, response.ThirdPartyReference)
	}

	if response.ResponseCode != "INS-10" {
		t.Fatalf("Transaction not processed successfully. Code: %s, Description: %s",
			response.ResponseCode, response.ResponseDescription)
	}
}
