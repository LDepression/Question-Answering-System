package settings

import (
	"wenba/internal/global"
	"wenba/internal/pkg/logger"
)

type log struct {
}

// Init 日志初始化
func (log) Init() {
	global.Logger = logger.NewLogger(&logger.InitStruct{
		LogSavePath:   global.Settings.Log.LogSavePath,
		LogFileExt:    global.Settings.Log.LogFileExt,
		MaxSize:       global.Settings.Log.MaxSize,
		MaxBackups:    global.Settings.Log.MaxBackups,
		MaxAge:        global.Settings.Log.MaxAge,
		Compress:      global.Settings.Log.Compress,
		LowLevelFile:  global.Settings.Log.LowLevelFile,
		HighLevelFile: global.Settings.Log.HighLevelFile,
	}, global.Settings.Log.Level)
}
