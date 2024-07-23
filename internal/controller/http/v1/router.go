package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"go-transaction-manager/internal/service"
	"log/slog"
)

func NewRouter(r *gin.Engine, services *service.Services, logger *slog.Logger) *gin.RouterGroup {
	v1 := r.Group("api/v1")
	{
		newUserRoutes(v1.Group("user"), services.User, logger)
		newReservationRoutes(v1.Group("reservation"), services.Reservation, logger)
		v1.GET("/swagger", swagger.WrapHandler(swaggerFiles.Handler))
	}

	return v1
}
