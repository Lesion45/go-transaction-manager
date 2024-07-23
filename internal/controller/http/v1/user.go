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

type userRoutes struct {
	userService service.User
	log         *slog.Logger
}

func newUserRoutes(g *gin.RouterGroup, userService service.User, log *slog.Logger) {
	r := &userRoutes{
		userService: userService,
		log:         log,
	}

	g.POST("/add_balance", r.addBalance())

	g.GET("/get_balance", r.getBalance())
}

type userAddBalanceInput struct {
	UserID uuid.UUID `json:"user-id" validate:"required"`
	Amount float64   `json:"amount" validate:"required"`
}

// @Summary		Add balance to user
// @Description	Returns status of operation
// @Tags		Balance
// @Accept		json
// @Produce		json
// @Param		input body v1.userAddBalanceInput true "input"
// @Success		200 {object} response.Response
// @Failure		400	{object} response.Response
// @Failure		500	{object} response.Response
// @Router		/api/v1/user/add_balance [post]
func (r *userRoutes) addBalance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "handlers.v1.User.addBalance"

		r.log.With(slog.String("op", op))

		var req userAddBalanceInput

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

		err = r.userService.AddBalance(ctx, service.UserAddBalanceInput{
			UserID: req.UserID,
			Amount: req.Amount,
		})
		if err != nil {
			r.log.Info("failed to add balance", sl.Err(err))
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("internal server error"))

			return
		}

		r.log.Info("balance added")

		ctx.IndentedJSON(http.StatusOK, response.OK())

		return
	}
}

type userGetBalanceInput struct {
	UserID uuid.UUID `json:"user-id" validate:"required"`
}

// @Summary		Get user balance
// @Description	Returns user balance
// @Tags		Balance
// @Accept		json
// @Produce		json
// @Param		input body v1.userGetBalanceInput true "input"
// @Success		200 {object} response.Balance
// @Failure		400	{object} response.Response
// @Failure		500	{object} response.Response
// @Router		/api/v1/user/get_balance [get]
func (r *userRoutes) getBalance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const op = "handlers.v1.User.getBalance"

		r.log.With(slog.String("op", op))

		var req userGetBalanceInput

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

		balance, err := r.userService.GetBalance(ctx, service.UserGetBalanceInput{
			UserID: req.UserID,
		})

		if err != nil {
			r.log.Info("failed to get balance")
			ctx.IndentedJSON(http.StatusInternalServerError, response.Error("internal server error"))

			return
		}

		r.log.Info("balance received")

		ctx.IndentedJSON(http.StatusOK, response.Balance{
			Balance: balance.Balance,
		})

		return
	}
}
