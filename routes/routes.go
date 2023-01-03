package routes

import (
	"net/http"

	"github.com/LouisMatos/challenge-backend-2-go/controller"
	"github.com/gin-gonic/gin"
)

func HandleRequest(Port string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.0.1"})

	apiReceitas := r.Group("/receitas")
	{
		apiReceitas.GET("/", controller.BuscaTodasReceitas)
		apiReceitas.GET("/:p1", controller.GetReceitaHandler)
		apiReceitas.GET("/:p1/:p2", controller.GetReceitaHandler)
		apiReceitas.POST("/", controller.CadastraReceita)
		apiReceitas.PUT("/:id", controller.AtualizarReceitaPorID)
		apiReceitas.DELETE("/:id", controller.DeletarReceitaPorID)
	}

	apiDespesa := r.Group("/despesas")
	{
		apiDespesa.GET("/", controller.BuscarTodasDespesas)
		apiDespesa.GET("/:p1", controller.GetDespesaHandler)
		apiDespesa.GET("/:p1/:p2", controller.GetDespesaHandler)
		apiDespesa.POST("/", controller.CadastrarDespesa)
		apiDespesa.PUT("/:id", controller.AtualizarDespesaPorID)
		apiDespesa.DELETE("/:id", controller.DeletarDespesaPorID)
	}

	r.GET("/healthcheck", controller.HealthCheck)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "Page not found"})
	})

	r.Run(":" + Port)
}
