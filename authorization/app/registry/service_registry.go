package registry

import (
	"go-klikdokter/app/service"

	"github.com/go-kit/log"
	"gorm.io/gorm"
)

func RegisterAuthService(db *gorm.DB, logger log.Logger) service.UserService {
	return service.NewUserService(
		logger,
	)
}
