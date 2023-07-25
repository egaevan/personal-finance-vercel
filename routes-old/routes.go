package routes_old

import (
	"github.com/gin-gonic/gin"
	"github.com/personal-finance-vercel/config"
	incomeHttp "github.com/personal-finance-vercel/delivery/http/income"
	incomeRepo "github.com/personal-finance-vercel/repository/pgsql/income"
	"net/http"
)

var (
	db         = config.InitPgsqlDB()
	repoIncome = incomeRepo.NewIncomeRepository(db)
	incomeHD   = incomeHttp.NewHandler(repoIncome)
)

func Register(route *gin.RouterGroup) {
	incomeRoute := route.Group("/income")
	{
		incomeRoute.GET("/:userId", incomeHD.GetIncome)
		incomeRoute.GET("/:userId/:id", incomeHD.FindIncome)
		incomeRoute.POST("/:userId/add", incomeHD.AddIncome)
		incomeRoute.PUT("/:userId/:id", incomeHD.PutIncome)
	}
}

func ErrRouter(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"errors": "this page could not be found",
	})
}
