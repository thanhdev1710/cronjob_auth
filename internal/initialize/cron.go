package initialize

import (
	"time"

	"github.com/robfig/cron/v3"
	"github.com/thanhdev1710/cronjob_auth/global"
	"github.com/thanhdev1710/cronjob_auth/internal/services"
	"go.uber.org/zap"
)

var CronJob *cron.Cron

func InitCronJobs() {
	// Khởi tạo cron theo định dạng chuẩn 5 trường (phút, giờ, ngày, tháng, thứ)
	CronJob = cron.New()

	// Chạy mỗi ngày lúc 2 giờ sáng
	_, err := CronJob.AddFunc("0 2 * * *", func() {
		now := time.Now().Format(time.RFC3339)

		global.Logger.Info("🌙 Bắt đầu xoá tài khoản hết hạn",
			zap.String("job", "delete_expired_accounts"),
			zap.String("time", now),
		)

		if err := services.NewUserService().DeleteExpiredAccounts(); err != nil {
			global.Logger.Error("❌ Xoá tài khoản thất bại", zap.Error(err))
		} else {
			global.Logger.Info("✅ Xoá tài khoản thành công", zap.String("time", now))
		}
	})

	if err != nil {
		global.Logger.Error("❌ CronJob init failed", zap.Error(err))
	}

	CronJob.Start()
}
