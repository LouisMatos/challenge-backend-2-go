package routes

import (
	"net/http"

	"github.com/LouisMatos/challenge-backend-2-go/app/controller"
	"github.com/LouisMatos/challenge-backend-2-go/app/middlewares"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

func HandleRequest(Port string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.SetTrustedProxies([]string{"192.168.0.1"})

	r.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())

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

	apiResumo := r.Group("/resumo")
	{
		apiResumo.GET("/:ano/:mes", controller.RealizarResumoAnoMes)
	}

	r.GET("/healthcheck", controller.HealthCheck)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "Page not found"})
	})

	r.Run(":" + Port)
}
