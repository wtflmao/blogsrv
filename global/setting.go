package global

import (
	"blogsrv/pkg/logger"
	settings "blogsrv/pkg/settings"
)

var (
	ServerSetting *settings.ServerSettingS
	AppSetting      *settings.AppSettingS
	DatabaseSetting *settings.DatabaseSettingS
)

var (
	Logger *logger.Logger
)
