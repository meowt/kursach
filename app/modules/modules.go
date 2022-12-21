package modules

import (
	"Diploma/pkg/theme"
	"Diploma/pkg/user"
	userDelegate "Diploma/pkg/user/delegate"
	userGateway "Diploma/pkg/user/gateway"
	userHttp "Diploma/pkg/user/handler"
	userUsecase "Diploma/pkg/user/usecase"

	themeDelegate "Diploma/pkg/theme/delegate"
	themeDrive "Diploma/pkg/theme/drive"
	themeGateway "Diploma/pkg/theme/gateway"
	themeHttp "Diploma/pkg/theme/handler"
	themeUsecase "Diploma/pkg/theme/usecase"

	"github.com/jmoiron/sqlx"
)

// HandlerModule section
type HandlerModule struct {
	UserHandler  userHttp.Handler
	ThemeHandler themeHttp.Handler
}

func SetupHandler(delegate DelegateModule) HandlerModule {
	return HandlerModule{
		UserHandler:  userHttp.SetupUserHandler(delegate.UserDelegate),
		ThemeHandler: themeHttp.SetupThemeHandler(delegate.ThemeDelegate),
	}
}

// DelegateModule section
type DelegateModule struct {
	UserDelegate  user.Delegate
	ThemeDelegate theme.Delegate
}

func SetupDelegate(usecaseModule UseCaseModule) DelegateModule {
	return DelegateModule{
		UserDelegate:  userDelegate.SetupUserDelegate(usecaseModule),
		ThemeDelegate: themeDelegate.SetupThemeDelegate(usecaseModule),
	}
}

// UseCaseModule section
type UseCaseModule struct {
	UserUseCase  user.UseCase
	ThemeUseCase theme.UseCase
}

func SetupUseCase(gatewayModule GatewayModule, driveModule DriveModule) UseCaseModule {
	return UseCaseModule{
		UserUseCase:  userUsecase.SetupUserUseCase(gatewayModule),
		ThemeUseCase: themeUsecase.SetupThemeUseCase(gatewayModule, driveModule),
	}
}

type DriveModule struct {
	ThemeDrive theme.Drive
}

func SetupDrive() DriveModule {
	return DriveModule{
		ThemeDrive: themeDrive.SetupThemeDrive(),
	}
}

// GatewayModule section
type GatewayModule struct {
	UserGateway  user.Gateway
	ThemeGateway theme.Gateway
}

func SetupGateway(PostgresClient *sqlx.DB) GatewayModule {
	return GatewayModule{
		UserGateway:  userGateway.SetupUserGateway(PostgresClient),
		ThemeGateway: themeGateway.SetupThemeGateway(PostgresClient),
	}
}
