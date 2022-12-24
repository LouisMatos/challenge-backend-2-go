package routes

import (
	"github.com/LouisMatos/challenge-backend-2-go/controller"
	"github.com/gin-gonic/gin"
)

func HandleRequest(Port string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.0.1"})
	r.GET("/healthcheck", controller.HealthCheck)
	r.POST("/receitas", controller.CadastraReceita)
	r.GET("/receitas", controller.BuscaTodasReceitas)
	r.Run(":" + Port)
}
