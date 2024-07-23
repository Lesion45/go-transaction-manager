package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go-transaction-manager/internal/controller/http/response"
	"go-transaction-manager/internal/service"
	"go-transaction-manager/pkg/logger/sl"
	"log/slog"
	"net/http"
)

type reservationRoutes struct {
	reservationService service.Reservation
	log                *slog.Logger
}

func newReservationRoutes(g *gin.RouterGroup, reservationService service.Reservation, log *slog.Logger) {
	r := &reservationRoutes{
		reservationService: reservationService,
		log:                log,
	}

	g.POST("/reserve_balance", r.reserveBalance())

	g.GET("/commit_balance", r.commitReservedBalance())
}

type reservationReserveBalance struct {
	OrderID   uuid.UUID `json:"order_id" validate:"required"`
	UserID    uuid.UUID `json:"user-id" validate:"required"`
	ServiceID uuid.UUID `json:"service-id" validate:"required"`
	Amount    float64   `json:"amount" validate:"required"`
	Info      string    `json:"info"`
}

// @Summary		Reserve user balance
// @Description	Returns status of operation
// @Tags		Balance
// @Accept		json
// @Produce		json
// @Param		input body v1.reservationReserveBalance true "input"
// @Success		200 {object} response.Response
// @Failure		400	{object} response.Response
// @Failure		500	{object} response.Response
// @Router		/api/v1/reservation/reserve_balance [post]
func (r *reservationRoutes) reserveBalance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "handlers.v1.Reservation.reserveBalance"

		r.log.With(slog.String("op", op))

		var req reservationReserveBalance

		err := ctx.BindJSON(&req)
		if err != nil {
			r.log.Error("failed to decode request body", sl.Err(err))
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("failed to decode request"))

			return
		}

		r.log.Info("request body decoded", slog.Any("request", req))
		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			r.log.Error("invalid request", sl.Err(err))
			ctx.IndentedJSON(http.StatusBadRequest, response.ValidataionError(validateErr))

			return
		}

		err = r.reservationService.ReserveBalance(ctx, service.ReservationReserveBalanceInput{
			OrderID:   req.OrderID,
			UserID:    req.UserID,
			ServiceID: req.ServiceID,
			Amount:    req.Amount,
			Info:      req.Info,
		})
		if err != nil {
			r.log.Info("failed to reserve balance")
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("internal server error"))

			return
		}

		r.log.Info("balance reserved")

		ctx.IndentedJSON(http.StatusOK, response.OK())

		return
	}
}

type reservationCommitReservedBalance struct {
	OrderID   uuid.UUID `json:"order_id" validate:"required"`
	UserID    uuid.UUID `json:"user-id" validate:"required"`
	ServiceID uuid.UUID `json:"service-id" validate:"required"`
	Amount    float64   `json:"amount" validate:"required"`
	Info      string    `json:"info"`
}

// @Summary		Commit reserved user balance
// @Description	Returns status of operation
// @Tags		Balance
// @Accept		json
// @Produce		json
// @Param		input body v1.reservationCommitReservedBalance true "input"
// @Success		200 {object} response.Response
// @Failure		400	{object} response.Response
// @Failure		500	{object} response.Response
// @Router		/api/v1/reservation/commit_balance [post]
func (r *reservationRoutes) commitReservedBalance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "handlers.v1.Reservation.commitReservedBalance"

		var req reservationCommitReservedBalance

		err := ctx.BindJSON(&req)
		if err != nil {
			r.log.Error("failed to decode request body", sl.Err(err))
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("failed to decode request"))

			return
		}

		r.log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			r.log.Error("invalid request", sl.Err(err))
			ctx.IndentedJSON(http.StatusBadRequest, response.ValidataionError(validateErr))

			return
		}

		err = r.reservationService.CommitReservedBalance(ctx, service.ReservationCommitReservedBalanceInput{
			OrderID:   req.OrderID,
			UserID:    req.UserID,
			ServiceID: req.ServiceID,
			Amount:    req.Amount,
			Info:      req.Info,
		})
		if err != nil {
			r.log.Info("failed to commit reserved balance")
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("internal server error"))

			return
		}

		r.log.Info("reserved balance committed")

		ctx.IndentedJSON(http.StatusOK, response.OK())

		return
	}
}
