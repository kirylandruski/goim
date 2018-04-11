package server

type Authenticator interface {
	Authenticate(username *string, password *string) (bool, error)
}
