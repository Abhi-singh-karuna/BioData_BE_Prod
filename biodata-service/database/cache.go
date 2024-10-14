package database

import (
	config "myapp/config"

	"github.com/Abhi-singh-karuna/my_Liberary/baselogger"
	"github.com/Abhi-singh-karuna/my_Liberary/cachehandler"
)

func ConnectRedis(cfg *config.Config, log *baselogger.BaseLogger) (cachehandler.CacheHandler, error) {
	log.Infof("Host: %v -- Port: %v -- Password: %v -- DB: %v", cfg.Redis.Write.Host, cfg.Redis.Write.Port, cfg.Redis.Write.Password, cfg.Redis.Write.Database)

	// Create Redis config
	redisConfig := cachehandler.Redis{
		Host:     cfg.Redis.Write.Host,
		Database: cfg.Redis.Write.Database,
		Password: cfg.Redis.Write.Password,
		Port:     cfg.Redis.Write.Port,
	}

	// Create Redis handler
	redisHandler := cachehandler.NewCacheHandler(redisConfig, log)

	return redisHandler, nil
}
