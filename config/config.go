package config

type Config struct {
	App   *App
	Slack *Slack
	Redis *Redis
}

func NewConfig() *Config {
	return &Config{
		App: &App{
			PokerRanks: getListStrEnv("POKER_RANKS", "?,0,0.5,1,2,3,5,8,13,20,40,100"),
		},
		Slack: &Slack{
			Token:         getStrEnv("SLACK_TOKEN", "xoxb-20686481348-1727532387907-XtSMmOxjZddDDh6N0JXi9o08"),
			ServerAddress: getStrEnv("SLACK_WEB_SERVER_ADDRESS", ":8000"),
		},
		Redis: &Redis{
			Host: getStrEnv("REDIS_HOST", "0.0.0.0"),
			Port: getIntEnv("REDIS_PORT", "6379"),
			DB:   getIntEnv("REDIS_DB", "0"),
		},
	}
}
