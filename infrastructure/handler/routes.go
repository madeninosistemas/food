package handler

import (
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/ninosistemas10/delivery/infrastructure/handler/category"
	"github.com/ninosistemas10/delivery/infrastructure/handler/login"
	"github.com/ninosistemas10/delivery/infrastructure/handler/mesa"
	"github.com/ninosistemas10/delivery/infrastructure/handler/producto"
	"github.com/ninosistemas10/delivery/infrastructure/handler/promocion"
	"github.com/ninosistemas10/delivery/infrastructure/handler/user"
)

func InitRoutes(e *echo.Echo, dbPool *pgxpool.Pool) {
	health(e)

	// A
	// B
	// C
	category.NewRouter(e, dbPool)
	// I


	// L
	login.NewRouter(e, dbPool)

	//M
	mesa.NewRouter(e, dbPool)

	// P
	producto.NewRouter(e, dbPool)
	promocion.NewRouter(e, dbPool)

	// U
	user.NewRouter(e, dbPool)
}

func health(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			map[string]string{
				"time":         time.Now().String(),
				"message":      "Hello World!",
				"service_name": "",
			},
		)
	})
}
