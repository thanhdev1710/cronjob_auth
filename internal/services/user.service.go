package services

import (
	"fmt"
	"time"

	"github.com/thanhdev1710/cronjob_auth/global"
	"github.com/thanhdev1710/cronjob_auth/internal/models"
	"go.uber.org/zap"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) DeleteExpiredAccounts() error {
	threeDaysAgo := time.Now().Add(-72 * time.Hour)

	var users []models.User
	err := global.Pdb.
		Where("(status = ? OR status = ?) AND deleted_at <= ?", global.User.Inactive, global.User.Banned, threeDaysAgo).
		Find(&users).Error

	if err != nil {
		return fmt.Errorf("failed to query expired accounts: %w", err)
	}

	if len(users) == 0 {
		global.Logger.Info("Không có tài khoản nào để xoá.")
		return nil
	}

	// Log từng user
	for _, user := range users {
		global.Logger.Info("🗑️ Sắp xoá user",
			zap.String("id", user.Id.String()),
			zap.String("email", user.Email),
			zap.Time("deleted_at", *user.DeletedAt),
		)
	}

	// Thực hiện xoá
	result := global.Pdb.Delete(users)

	if result.Error != nil {
		return fmt.Errorf("failed to delete expired accounts: %w", result.Error)
	}

	global.Logger.Info("✅ Đã xoá user hết hạn",
		zap.Int64("rows_affected", result.RowsAffected),
	)

	return nil
}
