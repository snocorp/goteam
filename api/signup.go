package api

type SignUpService interface {
	SignUp(name, email, password string) (u User, s Session, err error)
}
