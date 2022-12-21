package delegate

import (
	"Diploma/pkg/user"
)

type UserDelegateImpl struct {
	user.UseCase
}

type UserGatewayModule struct {
	user.Gateway
}

func SetupUserDelegate(usecase user.UseCase) UserGatewayModule {
	return UserGatewayModule{
		Gateway: &UserDelegateImpl{UseCase: usecase},
	}
}
