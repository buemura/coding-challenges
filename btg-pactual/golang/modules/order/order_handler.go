package order

// func SetupRoutes(e *echo.Echo) {
// 	e.GET("/order", getOrderTotalAmount)
// 	e.GET("/order/:orderId/total-amount", getOrderTotalAmount)
// }

// func getOrderTotalAmount(c echo.Context) error {
// 	orderIdStr := c.Param("orderId")
// 	orderId, err := strconv.Atoi(orderIdStr)
// 	if err != nil {
// 		return c.NoContent(http.StatusUnprocessableEntity)
// 	}

// 	oSevice := NewOrderService()

// 	stt, err := oSevice.getCustomerOrderList(o)
// 	if err != nil {
// 		return helper.HandleHttpError(c, err)
// 	}

// 	return c.JSON(http.StatusOK, stt)
// }
