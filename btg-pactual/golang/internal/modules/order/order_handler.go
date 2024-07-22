package order

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/customer/:customerId/orders", getCustomerOrderList)
}

func getCustomerOrderList(c echo.Context) error {
	customerIdStr := c.Param("customerId")
	customerId, err := strconv.Atoi(customerIdStr)
	if err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	var page, items int = 1, 5
	pageStr := c.QueryParam("page")
	if len(pageStr) > 1 {
		page, _ = strconv.Atoi(pageStr)
	}
	itemsStr := c.QueryParam("items")
	if len(itemsStr) > 1 {
		items, _ = strconv.Atoi(itemsStr)
	}

	oSevice := NewOrderService()
	res, err := oSevice.getCustomerOrderList(&OrderListIn{
		CustomerID: customerId,
		Page:       page,
		Items:      items,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
