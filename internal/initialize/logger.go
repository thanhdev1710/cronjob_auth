package initialize

import (
	"github.com/thanhdev1710/cronjob_auth/global"
	"github.com/thanhdev1710/cronjob_auth/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
