package config

type Config struct {
	App   *App
	Slack *Slack
	Redis *Redis
}

func NewConfig() *Config {
	return &Config{
		App: &App{
			ServerAddress: getStrEnv("WEB_SERVER_ADDRESS", ":8000"),
			PokerRanks:    getListStrEnv("POKER_RANKS", "?,0,0.5,1,2,3,5,8,13,20,40,100"),
		},
		Slack: &Slack{
			Token: getStrEnv("SLACK_TOKEN", "FILL_ME"),
		},
		Redis: &Redis{
			Host: getStrEnv("REDIS_HOST", "0.0.0.0"),
			Port: getIntEnv("REDIS_PORT", "6379"),
			DB:   getIntEnv("REDIS_DB", "0"),
		},
	}
}
