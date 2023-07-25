package api

import (
	"github.com/gin-gonic/gin"
	"github.com/personal-finance-vercel/config"
	incomeHttp "github.com/personal-finance-vercel/delivery/http/income"
	incomeRepo "github.com/personal-finance-vercel/repository/pgsql/income"
	"net/http"
)

var (
	app        *gin.Engine
	basePath   = "/api/serv1"
	db         = config.InitPgsqlDB()
	repoIncome = incomeRepo.NewIncomeRepository(db)
	incomeHD   = incomeHttp.NewHandler(repoIncome)
)

func init() {

	app = gin.Default()
	app.NoRoute(ErrRouter)
	route := app.Group(basePath)

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

// Entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
