package response

import "github.com/golang/protobuf/ptypes/timestamp"

type LoginResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
	User   `json:"user"`
}

type LoginByGoogleResponse struct {
	AccessToken    string              `json:"access_token"`
	RefreshToken   string              `json:"refresh_token"`
	ExpirationTime timestamp.Timestamp `json:"expiration_time"`
}
type User struct {
	Id                    string              `json:"id"`
	Name                  string              `json:"name"`
	Email                 string              `json:"email"`
	Phone                 string              `json:"phone"`
	Mobile                int                 `json:"mobile"`
	Address               string              `json:"address"`
	Role                  string              `json:"role"`
	Default_checkin_time  string              `json:"default_checkin_time"`
	Default_checkout_time string              `json:"default_checkout_time"`
	Title                 string              `json:"title"`
	In_report             string              `json:"in_report"`
	Is_deleted            string              `json:"is_deleted"`
	Holidays              string              `json:"holidays"`
	User_manager          string              `json:"user_manager"`
	Order_number          string              `json:"order_number"`
	Project_join          string              `json:"project_join"`
	Birthday              timestamp.Timestamp `json:"birthday"`
	Avatar                string              `json:"avatar"`
	Create_date           string              `json:"create_date"`
	Is_time               bool                `json:"is_time"`
	Start_date            timestamp.Timestamp `json:"start_date"`
	End_date              timestamp.Timestamp `json:"end_date"`
	Cover                 int                 `json:"cover"`
	Push_token            string              `json:"push_token"`
	Days_off              string              `json:"days_off"`
	Is_locked             bool                `json:"is_locked"`
	Department            string              `json:"department"`
	Data_access           string              `json:"data_access"`
	Probationary_day      int                 `json:"probationary_day"`
	Division              string              `json:"division"`
}
