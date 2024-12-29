package bootstrap

import (
	"app/internal/business"
	"app/internal/business/usecases/auth"
	"app/internal/business/usecases/data"
	"app/internal/infra/postgres"
	"context"
)

func InitAppDeps() (*AppDeps, error) {
	cfg, err := ReadAppConfig(ReadEnv())
	if err != nil {
		return nil, err
	}

	postgresRepo, err := postgres.InitPostgresRepo(cfg.DbConnStr, context.Background())
	if err != nil {
		return nil, err
	}

	bisDeps := &business.BusinessDeps{
		UserRepo: postgresRepo,
	}

	authUseCases := auth.NewAuthUseCases(bisDeps)
	dataUseCases := data.NewDataUseCases(bisDeps)

	return &AppDeps{
		AuthUseCases: authUseCases,
		DataUseCases: dataUseCases,
		Cfg:          cfg,
	}, nil
}
