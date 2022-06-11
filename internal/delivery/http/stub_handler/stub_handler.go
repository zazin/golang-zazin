package stub_handler

import (
	"github.com/gorilla/mux"
)

type stubHandler struct {
}

func StubHandler(r *mux.Router) {
	handler := &stubHandler{}

	r.HandleFunc("/stub/payment_va", handler.SubmitPayment).Methods("POST")
}
