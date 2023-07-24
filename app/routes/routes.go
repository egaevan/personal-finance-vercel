package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/personal-finance-vercel/app/internal/config"
	incomeHttp "github.com/personal-finance-vercel/app/internal/delivery/http/income"
	incomeRepo "github.com/personal-finance-vercel/app/internal/repository/pgsql/income"
	"net/http"
)

var (
	cfg        = config.Server()
	db         = config.InitPgsqlDB()
	repoIncome = incomeRepo.NewIncomeRepository(db)
	incomeHD   = incomeHttp.NewHandler(cfg, repoIncome)
)

func Register(app *gin.Engine) {
	app.NoRoute(ErrRouter)

	route := app.Group("/api")
	{
		incomeRoute := route.Group("/income")
		{
			incomeRoute.GET("/:userId", incomeHD.GetIncome)
			incomeRoute.GET("/:userId/:id", incomeHD.FindIncome)
			incomeRoute.POST("/:userId/add", incomeHD.AddIncome)
			incomeRoute.PUT("/:userId/:id", incomeHD.PutIncome)
		}
	}
}

func ErrRouter(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"errors": "this page could not be found brother",
	})
}
