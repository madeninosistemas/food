package producto

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/ninosistemas10/delivery/infrastructure/handler/middle"

	"github.com/ninosistemas10/delivery/domain/producto"
	productoStorage "github.com/ninosistemas10/delivery/infrastructure/postgres/producto"
)

func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	authMiddleware := middle.New()

	adminRoutes(e, h, authMiddleware.IsValid, authMiddleware.IsAdmin)
	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCase := producto.New(productoStorage.New(dbPool))
	return newHandler(useCase)
}

func adminRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := e.Group("/ninosistemas/admin/productos", middlewares...)

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)
	route.PUT("/upDateNoImagen/:id", h.UpdateEsceptImage)
	route.GET("", h.GetAll)
	//route.GET("/:id", h.GetByID)
}

func publicRoutes(e *echo.Echo, h handler) {
	route := e.Group("/ninosistemas/public/productos")

	route.POST("", h.Create)
	route.GET("", h.GetAll)
	route.GET("/:id", h.GetByID)
	route.GET("/categoria/:idcategoria", h.GetByCategoryID)
}
