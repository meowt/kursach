package usecase

import (
	"Diploma/pkg/user"
)

type UserUseCaseImpl struct {
	user.Gateway
}

type UserUseCaseModule struct {
	user.UseCase
}

func SetupUserUseCase(gateway user.Gateway) UserUseCaseModule {
	return UserUseCaseModule{
		UseCase: &UserUseCaseImpl{Gateway: gateway},
	}
}
