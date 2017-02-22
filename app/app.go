package app

import (
	"github.com/desertbit/glue"

	"github.com/nazarnovak/mind2/config"
	"github.com/nazarnovak/mind2/postgres"
	"github.com/nazarnovak/mind2/redis"
	"github.com/nazarnovak/mind2/session"
)

type App struct {
	Config *config.Config
	DB *postgres.DB
	Glue *glue.Server
	Redis *redis.Redis
	Session *session.Session
}

func (a *App) Init() error {
	a.Config = &config.Config{}
	a.DB = &postgres.DB{}
	a.Redis = &redis.Redis{}
	a.Session = &session.Session{}

	err := a.Config.Load()
	if err != nil {
		return err
	}

	err = a.DB.Open(
		a.Config.DB.Host,
		a.Config.DB.Port,
		a.Config.DB.User,
		a.Config.DB.Password,
		a.Config.DB.Name,
	)
	if err != nil {
		return err
	}

	err = a.Redis.Init(a.Config.RedisURL)
	if err != nil {
		return err
	}

	a.Session.Init()

	a.Glue = glue.NewServer(glue.Options{
		HTTPSocketType: glue.HTTPSocketTypeNone,
	})
	a.Glue.OnNewSocket(a.Redis.HandleSocket)

	return nil
}
