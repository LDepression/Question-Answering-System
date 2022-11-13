package settings

import (
	"wenba/internal/global"
	"wenba/internal/pkg/snowflake"
)

type sf struct {
}

// Init 雪花算法初始化
func (sf) Init() {
	var err error
	if global.Snowflake, err = snowflake.Init(global.Settings.App.StartTime, 1); err != nil {
		panic(err)
	}
}
