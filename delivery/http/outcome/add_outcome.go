package income

import (
	"github.com/gin-gonic/gin"
	"github.com/personal-finance-vercel/domain/entity"
	"github.com/personal-finance-vercel/pkg/utils"
	"net/http"
)

func (h *outcomeHandler) AddOutcome(ctx *gin.Context) {

	dataReq := entity.Outcome{}

	reqId := ctx.Param("userId")
	userId := utils.ConvertStrToInt(reqId)

	if err := ctx.Bind(&dataReq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.RestBody{
			Message: "invalid data request",
		})
		return
	}

	dataReq.UserId = userId

	err := h.outcomeUseCase.AddOutcome(ctx, &dataReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.RestBody{
			Message: "internal error",
		})

		return
	}
	ctx.JSON(http.StatusCreated, utils.RestBody{
		Message: "success create",
	})
	return
}
