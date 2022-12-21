package usecase

import "Diploma/pkg/theme"

type ThemeUseCaseImpl struct {
	theme.Drive
	theme.Gateway
}

type ThemeUseCaseModule struct {
	theme.UseCase
}

func SetupThemeUseCase(gateway theme.Gateway, drive theme.Drive) ThemeUseCaseModule {
	return ThemeUseCaseModule{
		UseCase: &ThemeUseCaseImpl{Gateway: gateway, Drive: drive},
	}
}
