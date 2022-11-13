package settings

import (
	"wenba/internal/dao"
	db "wenba/internal/dao/mysql"
	"wenba/internal/dao/redis"
	"wenba/internal/global"
)

type mDao struct {
}

// Init 持久化层初始化
func (m mDao) Init() {
	dao.Group.DB = db.Init(global.Settings.MySQL.User, global.Settings.MySQL.Password, global.Settings.MySQL.Host, global.Settings.MySQL.Port, global.Settings.MySQL.DbName)
	dao.Group.Redis = redis.Init(global.Settings.Redis.Address, global.Settings.Redis.Password, global.Settings.Redis.PoolSize, global.Settings.Redis.DB)
}
