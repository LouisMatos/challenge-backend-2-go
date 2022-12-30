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
	r.GET("/receitas", controller.BuscaTodasReceitas)
	r.GET("/receitas/:id", controller.BuscarReceitaId)
	r.GET("/despesas/:id", controller.BuscarDespesaId)
	r.GET("/despesas", controller.BuscarTodasDespesas)
	r.POST("/receitas", controller.CadastraReceita)
	r.POST("/despesas", controller.CadastrarDespesa)
	r.PUT("/receitas/:id", controller.AtualizarReceitaPorID)
	r.PUT("/despesas/:id", controller.AtualizarDespesaPorID)
	r.DELETE("/receitas/:id", controller.DeletarReceitaPorID)
	r.DELETE("/despesas/:id", controller.DeletarDespesaPorID)
	r.Run(":" + Port)
}
