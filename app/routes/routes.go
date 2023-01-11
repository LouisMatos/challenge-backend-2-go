package routes

import (
	"net/http"

	"github.com/LouisMatos/challenge-backend-2-go/app/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequest(Port string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE", "GET"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

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
