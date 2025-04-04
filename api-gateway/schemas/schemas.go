package schemas

import "fmt"

type Response struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description,omitempty"` // ok == false
	Result      any    `json:"result,omitempty"`      // ok == true
}

type RegisterRequest struct {
	FirstName      string `json:"first_name"`
	MiddleName     string `json:"middle_name"`
	LastName       string `json:"last_name"`
	BornDate       int32  `json:"born_date"`
	PhoneNumber    string `json:"phone_number"`
	Email          string `json:"email"`
	Password       string
	Specialization string
}

func (r RegisterRequest) Validate() error {
	// TODO
	return nil
}

type RegisterResponse struct {
	UserID   int32  `json:"user_id"`
	NewLogin string `json:"new_login"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

func (l LoginRequest) Validate() error {
	if l.Name == "" {
		return fmt.Errorf("name can't be empty - saying from gateway")
	}
	if len(l.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	return nil
}

type LoginResponse struct {
	Name string `json:"name"`
	ID   int32  `json:"id"`
	Role string `json:"role"`
}
