package transport

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"go-klikdokter/app/api/endpoint"
	"go-klikdokter/app/model/base/encoder"
	"go-klikdokter/app/model/request"
	"go-klikdokter/app/service"
	"net/http"
)

func UserHttpHandler(s service.UserService, logger log.Logger) http.Handler {
	us := mux.NewRouter()

	ep := endpoint.MakeUserEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		//httptransport.ServerErrorEncoder(encoder.EncodeError),
	}

	us.Methods("POST").Path("/users/login").Handler(httptransport.NewServer(
		ep.LoginEmail,
		decodeLoginEmail,
		encoder.EncodeResponseHTTP,
		options...,
	))

	return us
}

func decodeLoginEmail(ctx context.Context, r *http.Request) (rqst interface{}, err error) {
	var req request.LoginEmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
