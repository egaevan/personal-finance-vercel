package income

import (
	"github.com/gin-gonic/gin"
	"github.com/personal-finance-vercel/pkg/utils"
	"net/http"
)

func (h *expenseHandler) GetExpense(ctx *gin.Context) {

	reqId := ctx.Param("userId")
	userId := utils.ConvertStrToInt(reqId)

	res, err := h.expenseUseCase.GetExpense(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.RestBody{
			Message: "internal error",
		})

		return
	}
	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
	return
}
