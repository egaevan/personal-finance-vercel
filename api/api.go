package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/personal-finance-vercel/config"
	expenseHttp "github.com/personal-finance-vercel/delivery/http/expense"
	incomeHttp "github.com/personal-finance-vercel/delivery/http/income"
	expenseRepo "github.com/personal-finance-vercel/repository/pgsql/expense"
	incomeRepo "github.com/personal-finance-vercel/repository/pgsql/income"
	"net/http"
)

var (
	app         *gin.Engine
	basePath    = "/api/serv1"
	db          = config.InitPgsqlDB()
	repoIncome  = incomeRepo.NewIncomeRepository(db)
	repoExpense = expenseRepo.NewExpenseRepository(db)
	incomeHD    = incomeHttp.NewHandler(repoIncome)
	expenseHD   = expenseHttp.NewHandler(repoExpense)
)

func init() {

	app = gin.Default()
	app.NoRoute(ErrRouter)
	route := app.Group(basePath)

	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	incomeRoute := route.Group("/income")
	{
		incomeRoute.GET("/:userId", incomeHD.GetIncome)
		incomeRoute.GET("/:userId/:id", incomeHD.FindIncome)
		incomeRoute.POST("/:userId/add", incomeHD.AddIncome)
		incomeRoute.PUT("/:userId/:id", incomeHD.PutIncome)
	}

	expenseRoute := route.Group("/expense")
	{
		expenseRoute.GET("/:userId", expenseHD.GetExpense)
		expenseRoute.GET("/:userId/:id", expenseHD.FindExpense)
		expenseRoute.POST("/:userId/add", expenseHD.AddExpense)
		expenseRoute.PUT("/:userId/:id", expenseHD.PutExpense)
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
