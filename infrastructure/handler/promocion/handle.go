package promocion

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ninosistemas10/delivery/domain/promocion"
	"github.com/ninosistemas10/delivery/infrastructure/handler/response"
	"github.com/ninosistemas10/delivery/model"
)


type handler struct{
	useCase promocion.UseCase
	response response.API
}

func newHandler(useCase promocion.UseCase) handler {
	return handler{useCase: useCase}
}

func(h handler) Create(c echo.Context) error {
	m := model.Promocion{}
	if err := c.Bind(&m); err != nil{
		return h.response.BindFailed(err)
	}

	if err := h.useCase.Create(&m); err != nil {
		return h.response.Error(c, "useCase.Create()", err)
	}

	return c.JSON(h.response.Created(m))
}

func(h handler) Update(c echo.Context) error {
	m := model.Promocion{}
	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	ID, err := uuid.Parse(c.Param("id"))
	if err != nil { return h.response.BindFailed(err) }

	m.ID = ID
	if err := h.useCase.Update(&m); err != nil{
		return h.response.Error(c, "h.useCase.Update()", err)
	}
	return c.JSON(h.response.Updated(m))
}

func (h handler) Delete(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil { return h.response.BindFailed(err) }

	err = h.useCase.Delete(ID)
	if err != nil { return h.response.Error(c, "useCase.Delete()", err) }

	return c.JSON(h.response.Deleted(nil))
}

func (h handler) GetByID(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil { h.response.Error(c, "uuid.Parse()", err) }

	promocionData, err := h.useCase.GetByID(ID)
	if err != nil { return h.response.Error(c, "useCase.GetByID", err) }

	return c.JSON(h.response.OK(promocionData))
}

func (h handler) GetAll(c echo.Context) error {
	promociones, err := h.useCase.GetAll()
	if err != nil { return h.response.Error(c, "useCase.GetAll", err) }

	return c.JSON(h.response.OK(promociones))
}
