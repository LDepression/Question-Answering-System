package dao

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type group struct {
	DB    *sqlx.DB
	Redis *redis.Client
}

var Group = new(group)
