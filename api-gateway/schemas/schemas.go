package schemas

import "fmt"

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
	Name string
	ID   int32
	Role string
}
