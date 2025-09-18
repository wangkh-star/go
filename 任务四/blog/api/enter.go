package api

import "blog/api/settings"

type ApiGroup struct {
	Settingsapi settings.Settingsapi
}

var ApiGroupApp = new(ApiGroup)
