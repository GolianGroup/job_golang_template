package app

import "jobs_golang_template/internal/producers"

func (a *application) InitRedis() producers.RedisClient {
	return producers.NewRedis(&a.config.Redis)
}
