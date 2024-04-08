package login

import "github.com/ninosistemas10/delivery/model"

type UseCase interface {
	Login(email, password, jwtSecretKey string) (model.User, string, error)
}

type UseCaseUser interface {
	Login(email, password string) (model.User, error)
}
