package routes

import (
	"net/http"

	"github.com/LouisMatos/challenge-backend-2-go/app/controller"
	"github.com/LouisMatos/challenge-backend-2-go/app/database"
	"github.com/LouisMatos/challenge-backend-2-go/app/middlewares"
	"github.com/LouisMatos/challenge-backend-2-go/app/repository"
	"github.com/LouisMatos/challenge-backend-2-go/app/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	dbConnection             *gorm.DB
	receitaRepository        repository.ReceitaRepository
	despesaRepository        repository.DespesaRepository
	resumoRepository         repository.ResumoRepository
	usuarioRepository        repository.UsuarioRepository
	receitaService           service.ReceitaService
	despesaService           service.DespesaService
	resumoService            service.ResumoService
	usuarioService           service.UsuarioService
	loginService             service.LoginService
	jwtService               service.JWTService
	receitaController        controller.ReceitaController
	despesaController        controller.DespesaController
	resumoController         controller.ResumoController
	usuarioController        controller.UsuarioController
	loginController          controller.LoginController
	defaultReceitaController controller.DefaultReceitaController
	defaultDespesaController controller.DefaultDespesaController
)

func HandleRequest(Port string) {

	dbConnection = database.GetConnection()

	receitaRepository = repository.NewReceitaRepository(dbConnection)
	despesaRepository = repository.NewDespesaRepository(dbConnection)
	resumoRepository = repository.NewResumoRepository(dbConnection)
	usuarioRepository = repository.NewUsuarioRepository(dbConnection)

	receitaService = service.NewReceitaService(receitaRepository)
	despesaService = service.NewDespesaService(despesaRepository)
	resumoService = service.NewResumoService(resumoRepository)
	usuarioService = service.NewUsuarioService(usuarioRepository)
	loginService = service.NewLoginService()
	jwtService = service.NewJWTService()

	receitaController = controller.NewReceitaController(receitaService)
	despesaController = controller.NewDespesaController(despesaService)
	resumoController = controller.NewResumoController(resumoService)
	usuarioController = controller.NewUsuarioController(usuarioService)
	loginController = controller.NewLoginController(usuarioService, jwtService)

	defaultReceitaController = controller.NewDefaultReceitaController(receitaService)
	defaultDespesaController = controller.NewDefaultDespesaController(despesaService)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.SetTrustedProxies([]string{"192.168.0.1"})

	// r.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())

	apiReceitas := r.Group("/receitas", middlewares.AuthorizeJWT())
	{
		apiReceitas.GET("/", receitaController.GetAll)
		apiReceitas.GET("/:p1", defaultReceitaController.GetReceitaHandler)
		apiReceitas.GET("/:p1/:p2", defaultReceitaController.GetReceitaHandler)
		apiReceitas.POST("/", receitaController.Save)
		apiReceitas.PUT("/:id", receitaController.Update)
		apiReceitas.DELETE("/:id", receitaController.Delete)
	}

	apiDespesa := r.Group("/despesas", middlewares.AuthorizeJWT())
	{
		apiDespesa.GET("/", despesaController.GetAll)
		apiDespesa.GET("/:p1", defaultDespesaController.GetDespesaHandler)
		apiDespesa.GET("/:p1/:p2", defaultDespesaController.GetDespesaHandler)
		apiDespesa.POST("/", despesaController.Save)
		apiDespesa.PUT("/:id", despesaController.Update)
		apiDespesa.DELETE("/:id", despesaController.Delete)
	}

	apiResumo := r.Group("/resumo")
	{
		apiResumo.GET("/:ano/:mes", resumoController.GetMonthSummary)
	}

	apiUsuario := r.Group("/usuario")
	{
		apiUsuario.POST("/registrar", usuarioController.Save)
	}

	// Login Endpoint: Authentication + Token creation
	r.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.Status(http.StatusUnauthorized)
		}
	})

	r.GET("/healthcheck", controller.HealthCheck)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "Page not found"})
	})

	r.Run(":" + Port)
}
