package delegate

import "Diploma/pkg/theme"

type ThemeDelegateImpl struct {
	theme.UseCase
}

type ThemeGatewayModule struct {
	theme.Gateway
}

func SetupThemeDelegate(usecase theme.UseCase) ThemeGatewayModule {
	return ThemeGatewayModule{
		Gateway: &ThemeDelegateImpl{UseCase: usecase},
	}
}
