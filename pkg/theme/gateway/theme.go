package gateway

import (
	"Diploma/pkg/theme"
	"github.com/jmoiron/sqlx"
)

type ThemeGatewayImpl struct {
	PostgresClient *sqlx.DB
}

type ThemeGatewayModule struct {
	theme.Gateway
}

func SetupThemeGateway(postgresClient *sqlx.DB) ThemeGatewayModule {
	return ThemeGatewayModule{
		Gateway: &ThemeGatewayImpl{PostgresClient: postgresClient},
	}
}
