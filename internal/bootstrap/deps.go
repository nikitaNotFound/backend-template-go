package bootstrap

import (
	"app/internal/business/usecases/data"
	"app/internal/business/usecases/auth"
)

type AppDeps struct {
	AuthUseCases *auth.AuthUseCases
	DataUseCases *data.DataUseCases

	Cfg *AppConfig
}
