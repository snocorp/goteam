package kernel

import (
	"github.com/snocorp/goteam/api"
)

type signUpService struct {
	sessionSvc api.SessionService
	userSvc    api.UserService
}

func (svc *signUpService) SignUp(name, email, password string) (u api.User, s api.Session, err error) {
	u, err = svc.userSvc.Create(name, email, password)
	if err != nil {
		return
	}

	s, err = svc.sessionSvc.Create(u.ID, password)
	if err != nil {
		return
	}

	// TODO: add user to root application

	return
}
