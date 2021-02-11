package userscontracts

import "context"

//User information
type User struct {
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

//AuthenticationRequest struct to receibe on authentication
type AuthenticationRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

//UserService service
type UserService interface {
	Save(ctx context.Context, user User) error
	Get(ctx context.Context, username string) (u User, err error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, username string) error
	Autenticate(ctx context.Context, user AuthenticationRequest) (User, error)
}
