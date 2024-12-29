package bootstrap

import (
	"app/internal/business/usecases/data"
	"app/internal/business/usecases/user"
)

type AppDeps struct {
	AuthUseCases *user.AuthUseCases
	DataUseCases *data.DataUseCases

	Cfg *AppConfig
}
