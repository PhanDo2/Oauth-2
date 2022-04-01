package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/log"
	"go-klikdokter/app/model/request"
	"go-klikdokter/app/model/response"
	"go-klikdokter/helper/message"
	"io/ioutil"
	"net/http"
	"strings"
)

type UserService interface {
	LoginByEmail(input request.LoginEmailRequest) (*response.LoginResponse, message.Message)
}

type userServiceImpl struct {
	logger log.Logger
}

func NewUserService(
	lg log.Logger) UserService {
	return &userServiceImpl{lg}
}

func (s *userServiceImpl) LoginByEmail(input request.LoginEmailRequest) (*response.LoginResponse, message.Message) {
	//logger := log.With(s.logger, "UserService", "Login")
	url := "https://timesheet.savvycom.vn/users/loginByApi"
	method := "POST"
	fmt.Println("Email", input.Email)
	fmt.Println("pass", input.Password)
	payload := strings.NewReader(`{
    "email": "` + input.Email + `",
    "password": "` + input.Password + `"
}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, message.FailedLogin
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "SMS2014=ffc35fa12f1b9019c32d00ea99b07000")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, message.FailedLogin
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, message.FailedLogin
	}
	data := string(body)
	var LoginJSON response.LoginResponse
	_ = json.Unmarshal([]byte(data), &LoginJSON)

	if LoginJSON.Role == "member" {
		return nil, message.MemberNotLogin

	}

	return &LoginJSON, message.SuccessMsg

}
