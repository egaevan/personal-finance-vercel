package income

import (
	"github.com/gin-gonic/gin"
	"github.com/personal-finance-vercel/pkg/utils"
	"net/http"
)

func (h *expenseHandler) FindExpense(ctx *gin.Context) {

	reqId := ctx.Param("userId")
	id := ctx.Param("id")
	userId := utils.ConvertStrToInt(reqId)
	expenseId := utils.ConvertStrToInt(id)

	res, err := h.expenseUseCase.FindExpense(ctx, userId, expenseId)
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
