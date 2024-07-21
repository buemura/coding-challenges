package order

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/customer/:customerId/orders", getCustomerOrderList)
	e.GET("/customer/:customerId/orders-count", getCustomerOrdersCount)
	e.GET("/order/:orderId/total-price", getOrderTotalPrice)
}

func getCustomerOrderList(c echo.Context) error {
	customerIdStr := c.Param("customerId")
	customerId, err := strconv.Atoi(customerIdStr)
	if err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	oSevice := NewOrderService()

	res, err := oSevice.getCustomerOrderList(customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func getCustomerOrdersCount(c echo.Context) error {
	customerIdStr := c.Param("customerId")
	customerId, err := strconv.Atoi(customerIdStr)
	if err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	oSevice := NewOrderService()

	res, err := oSevice.getCustomerOrdersCount(customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func getOrderTotalPrice(c echo.Context) error {
	orderIdStr := c.Param("orderId")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	oSevice := NewOrderService()

	res, err := oSevice.getOrderTotalPrice(orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
