package stub_handler

import (
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"log"
	"net/http"
)

type jsonResponse struct {
	StatusCode      string `json:"status_code"`
	StatusDesc      string `json:"status_desc"`
	PaymentProvider string `json:"payment_provider"`
	PaymentResponse struct {
		ResponseCode string `json:"response_code"`
		ResponseDesc string `json:"response_desc"`
		DeeplinkUrl  string `json:"deeplink_url"`
		CheckoutUrl  string `json:"checkout_url"`
		QrUrl        string `json:"qr_url"`
		PaymentCode  string `json:"payment_code"`
	} `json:"payment_response"`
	Id string `json:"id"`
}

func (h *stubHandler) SubmitPayment(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(jsonResponse{
		StatusCode:      "00",
		StatusDesc:      "Success",
		PaymentProvider: "virtual_account",
		PaymentResponse: struct {
			ResponseCode string `json:"response_code"`
			ResponseDesc string `json:"response_desc"`
			DeeplinkUrl  string `json:"deeplink_url"`
			CheckoutUrl  string `json:"checkout_url"`
			QrUrl        string `json:"qr_url"`
			PaymentCode  string `json:"payment_code"`
		}{
			ResponseCode: "201",
			ResponseDesc: "Success",
			DeeplinkUrl:  "",
			CheckoutUrl:  "",
			QrUrl:        "",
			PaymentCode:  fmt.Sprintf("100000%v", faker.RandomUnixTime()),
		},
		Id: faker.UUIDDigit(),
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	write, err := w.Write(response)
	if err != nil {
		log.Println(write)
	}
}
