package settings

import (
	"wenba/internal/global"
	"wenba/internal/pkg/app"
)

type page struct {
}

func (page) Init() {
	global.Page = app.InitPage(global.Settings.Page.DefaultPageSize, global.Settings.Page.MaxPageSize,
		global.Page.PageKey, global.Page.PageSizeKey)

}
