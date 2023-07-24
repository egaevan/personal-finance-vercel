package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/personal-finance-vercel/app/handlers"
	"github.com/personal-finance-vercel/app/handlers/income_handlers"
	"github.com/personal-finance-vercel/app/internal/config"
	incomeRepo "github.com/personal-finance-vercel/app/internal/repository/pgsql/income"
	"net/http"
)

var (
	cfg        = config.Server()
	db         = config.InitPgsqlDB()
	repoIncome = incomeRepo.NewIncomeRepository(db)
	incomeHD   = income_handlers.NewHandler(cfg, repoIncome)
)

func Register(app *gin.Engine) {
	app.NoRoute(ErrRouter)

	app.GET("/ping", handlers.Ping)

	route := app.Group("/api")
	{
		test := route.Group("/test")
		{
			test.GET("/hello/:name", handlers.Hello)
		}
		income := route.Group("/income")
		{
			income.GET("/:userId", incomeHD.GetIncome)
		}
	}
}

func ErrRouter(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"errors": "this page could not be found",
	})
}
