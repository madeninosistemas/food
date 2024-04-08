package producto

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ninosistemas10/delivery/domain/producto"
	"github.com/ninosistemas10/delivery/model"

	"github.com/ninosistemas10/delivery/infrastructure/handler/response"
)

type handler struct {
	useCase producto.UseCase
	response response.API
}

func newHandler(useCase producto.UseCase) handler {
	return handler{useCase: useCase}
}

func (h handler) Create(c echo.Context) error {
    m := model.Producto{}

    if err := c.Bind(&m); err != nil {
        return h.response.BindFailed(err)
    }


    if err := h.useCase.Create(&m); err != nil {
        return h.response.Error(c, "useCase.Create()", err)
    }

    return c.JSON(h.response.Created(m))
}





func (h handler) Update(c echo.Context) error {
	m := model.Producto{}

	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.BindFailed(err)
	}

	m.ID = ID

	if err := h.useCase.Update(&m); err != nil {
		return h.response.Error(c,"h.useCase.Update()", err)
	}

	return c.JSON(h.response.Updated(m))
}

func (h handler) UpdateEsceptImage(c echo.Context) error {
	//bind de los datos del cuerpo de la solicitud a una instancia de model.Producto
	updatedProducto := model.Producto{}
	if err := c.Bind(&updatedProducto); err != nil {
		return h.response.BindFailed(err)
	}

	//Parsear el Id del producto de la URL
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.BindFailed(err)
	}

	//llamar al metodo UpdateEsceptImage de la useCase
	err = h.useCase.UpdateEsceptImage(ID, updatedProducto)
	if err != nil {
		return h.response.Error(c, "h.useCaase.Update()", err)
	}

	//retorna la repuesta JSON con el producto actualizadodd
	return c.JSON(h.response.Updated(updatedProducto))

}

func (h handler) Delete(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.BindFailed(err)
	}

	err = h.useCase.Delete(ID)
	if err != nil {
		return h.response.Error(c, "useCase.Delete()", err)
	}

	return c.JSON(h.response.Deleted(nil))
}

func (h handler) GetByID(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.Error(c, "uuid.Parse()", err)
	}

	productoData, err := h.useCase.GetByID(ID)
	if err != nil {
		return h.response.Error(c, "useCase.GetWhere()", err)
	}

	return c.JSON(h.response.OK(productoData))
}

func (h handler) GetByCategoryID(c echo.Context) error {
    idCategoria, err := uuid.Parse(c.Param("idcategoria"))
    fmt.Println("Valor importante de idcategoria:", idCategoria)

    // Imprime el contexto
    fmt.Printf("Contexto: %+v\n", c)

    if err != nil {
        return h.response.Error(c, "uuid.Parse()", err)
    }

    productos, err := h.useCase.GetByCategoryID(idCategoria)
    if err != nil {
        return h.response.Error(c, "useCase.GetByCategoryID", err)
    }

    return c.JSON(h.response.OK(productos))
}


func (h handler) GetAll(c echo.Context) error {
	productos, err := h.useCase.GetAll()
	if err != nil {
		return h.response.Error(c, "useCase.GetAllWhere()", err)
	}

	return c.JSON(h.response.OK(productos))
}

