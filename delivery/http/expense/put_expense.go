package income

import (
	"github.com/gin-gonic/gin"
	"github.com/personal-finance-vercel/domain/entity"
	"github.com/personal-finance-vercel/pkg/utils"
	"net/http"
)

func (h *expenseHandler) PutExpense(ctx *gin.Context) {

	reqId := ctx.Param("userId")
	id := ctx.Param("id")
	userId := utils.ConvertStrToInt(reqId)
	expenseId := utils.ConvertStrToInt(id)

	dataReq := entity.Expense{}

	if err := ctx.Bind(&dataReq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.RestBody{
			Message: "invalid data request",
		})
		return
	}

	err := h.expenseUseCase.PutExpenseById(ctx, &dataReq, userId, expenseId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.RestBody{
			Message: "internal error",
		})

		return
	}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
	})
	return
}
