package schemas

import "fmt"

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (l LoginRequest) Validate() error {
	if l.Name == "" {
		return fmt.Errorf("name can't be empty")
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
