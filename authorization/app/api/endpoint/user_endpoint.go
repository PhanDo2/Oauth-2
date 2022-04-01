package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go-klikdokter/app/model/base"
	"go-klikdokter/app/model/request"
	"go-klikdokter/app/service"
)

type UserEndpoint struct {
	LoginEmail endpoint.Endpoint
}

func MakeUserEndpoints(s service.UserService) UserEndpoint {
	return UserEndpoint{
		LoginEmail: makeLoginEmail(s),
	}
}

func makeLoginEmail(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (resp interface{}, err error) {
		req := rqst.(request.LoginEmailRequest)
		result, msg := s.LoginByEmail(req)
		if msg.Code == 4000 {
			return base.SetHttpResponse(msg.Code, msg.Message, nil, nil), nil
		}
		return base.SetHttpResponse(msg.Code, msg.Message, result, nil), nil
	}
}
