package category

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ninosistemas10/delivery/domain/category"
	"github.com/ninosistemas10/delivery/infrastructure/handler/response"
	"github.com/ninosistemas10/delivery/model"
)



type handler struct {
	useCase category.UseCase
	response response.API
}

func newHandler(useCase category.UseCase) handler {
	return handler{useCase: useCase}
}

func (h handler) Create(c echo.Context) error {
	m := model.Category{}
	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	if err := h.useCase.Create(&m); err != nil {
		return h.response.Error(c, "useCase.Create()",err) }

	return c.JSON(h.response.Created(m))
}


func (h handler) Update(c echo.Context) error {
	m := model.Category{}
	if err := c.Bind(&m); err != nil { return h.response.BindFailed(err) }

	ID, err := uuid.Parse(c.Param("id"))
	if err != nil { return h.response.BindFailed(err)}

	m.ID = ID

	if err := h.useCase.Update(&m); err != nil {
		return h.response.Error(c,"h.useCase.Update()", err)
	}

	return c.JSON(h.response.Updated(m))
}

func (h handler) UpdateEsceptImage(c echo.Context) error {
    // Bind de los datos del cuerpo de la solicitud a una instancia de model.Category
    updatedCategory := model.Category{}
    if err := c.Bind(&updatedCategory); err != nil {
        return h.response.BindFailed(err)
    }

    // Parsear el ID de la categoría de la URL
    ID, err := uuid.Parse(c.Param("id"))
    if err != nil {
        return h.response.BindFailed(err)
    }

    // Llamar al método UpdateEsceptImage de la useCase
    err = h.useCase.UpdateEsceptImage(ID, updatedCategory)
    if err != nil {
        return h.response.Error(c, "h.useCase.UpdateEsceptImage()", err)
    }

    // Retornar la respuesta JSON con la categoría actualizada
    return c.JSON(h.response.Updated(updatedCategory))
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

	categoryData, err := h.useCase.GetByID(ID)
	if err != nil { return h.response.Error(c,"useCase.GetBYID",err) }

	return c.JSON(h.response.OK(categoryData))
}

func (h handler) GetAll(c echo.Context) error {
	categorys, err := h.useCase.GetAll()
	if err != nil { return h.response.Error(c, "useCase.GetAll", err) }

	return c.JSON(h.response.OK(categorys))
}

