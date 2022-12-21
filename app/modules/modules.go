package modules

import (
	"Diploma/pkg/user"
	userDelegate "Diploma/pkg/user/delegate"
	userGateway "Diploma/pkg/user/gateway"
	userHttp "Diploma/pkg/user/handler"
	userUsecase "Diploma/pkg/user/usecase"
	"github.com/jmoiron/sqlx"
)

// HandlerModule section
type HandlerModule struct {
	UserHandler userHttp.Handler
}

func SetupHandler(delegate DelegateModule) HandlerModule {
	return HandlerModule{
		UserHandler: userHttp.SetupUserHandler(delegate.UserDelegate),
	}
}

// DelegateModule section
type DelegateModule struct {
	UserDelegate user.Delegate
}

func SetupDelegate(usecaseModule UseCaseModule) DelegateModule {
	return DelegateModule{
		UserDelegate: userDelegate.SetupUserDelegate(usecaseModule),
	}
}

// UseCaseModule section
type UseCaseModule struct {
	UserUseCase user.UseCase
}

func SetupUseCase(gatewayModule GatewayModule) UseCaseModule {
	return UseCaseModule{
		UserUseCase: userUsecase.SetupUserUseCase(gatewayModule),
	}
}

// GatewayModule section
type GatewayModule struct {
	UserGateway user.Gateway
}

func SetupGateway(PostgresClient *sqlx.DB) GatewayModule {
	return GatewayModule{
		UserGateway: userGateway.SetupUserGateway(PostgresClient),
	}
}
