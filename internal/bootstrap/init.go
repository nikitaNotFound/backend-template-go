package bootstrap

func InitAppDeps() (*AppDeps, error) {
	cfg, err := ReadAppConfig(ReadEnv())
	if err != nil {
		return nil, err
	}

	return &AppDeps{
		Cfg: cfg,
	}, nil
}
