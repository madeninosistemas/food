package mesa

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/ninosistemas10/delivery/domain/mesa"
	"github.com/ninosistemas10/delivery/infrastructure/handler/middle"
	mesaStorage "github.com/ninosistemas10/delivery/infrastructure/postgres/mesa"
)


func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := builHandler(dbPool)

	authMiddlleware := middle.New()

	adminRoutes(e, h, authMiddlleware.IsValid, authMiddlleware.IsAdmin)
	publicRoutes(e, h)
}

func builHandler(dbPool * pgxpool.Pool) handler {
	useCase := mesa.New(mesaStorage.New(dbPool))

	return newHandler(useCase)
}

func adminRoutes(e * echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := e.Group("ninosistemas/admin/mesa", middlewares...)

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)

	route.GET("", h.GetAll)
	route.GET("/:id",h.GetByID)
}

func publicRoutes(e * echo.Echo, h handler) {
	route := e.Group("ninosistemas/public/mesa")

	route.POST("", h.Create)
	route.GET("", h.GetAll)
}
