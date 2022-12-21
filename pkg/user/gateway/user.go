package gateway

import (
	"Diploma/pkg/user"
	"github.com/jmoiron/sqlx"
)

type UserGatewayImpl struct {
	PostgresClient *sqlx.DB
}

type UserGatewayModule struct {
	user.Gateway
}

func SetupUserGateway(postgresClient *sqlx.DB) UserGatewayModule {
	return UserGatewayModule{
		Gateway: &UserGatewayImpl{PostgresClient: postgresClient},
	}
}
