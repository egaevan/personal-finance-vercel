package api_old

import (
	"github.com/gin-gonic/gin"
	"github.com/personal-finance-vercel/routes-old"
	"net/http"
)

var (
	app      *gin.Engine
	basePath = "/api"
)

// @title Golang Vercel Deployment
// @description API Documentation for Golang deployment in vercel serverless environment
// @version 1.0

// @schemes https http
// @host golang-vercel.vercel.app
func init() {

	app = gin.Default()
	app.NoRoute(routes_old.ErrRouter)
	basePath := app.Group(basePath)
	routes_old.Register(basePath)
}

// Entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
