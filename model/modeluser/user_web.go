package modeluser

import "time"

type Request struct {
	ID       uint   `json:"id,omitempty"`
	Username string `json:"username" example:"jhondoe"`
	Email    string `json:"email" example:"test@example.com"`
	Password string `json:"password" example:"password"`
	Age      int    `json:"age" example:"23"`
}

type Response struct {
	ID        uint      `json:"id"  example:"1"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2021-11-03T01:52:41.035Z"`
	Username  string    `json:"username"  example:"jhondoe"`
	Email     string    `json:"email" example:"test@example.com"`
	Age       int       `json:"age" example:"23"`
}

type RequestLogin struct {
	Email    string `json:"email" example:"test@example.com"`
	Password string `json:"password" example:"password"`
}

type ResponseLogin struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJxd2Vxd2..."`
}

// ExampleRequestUpdate only for example swaggo docs
type ExampleRequestUpdate struct {
	Username string `json:"username" example:"jhondoe"`
	Email    string `json:"email" example:"test@example.com"`
}

type ExampleResponseDelete struct {
	Message string `json:"message" example:"your account has been successfully deleted"`
}
