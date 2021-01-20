package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func getMessages(c echo.Context) (err error) {
	if messages, count := dbGetMessages(); count > 0 {
		n := MESS{messages}
		err = c.JSON(http.StatusOK, n)
	} else {
		err = c.NoContent(http.StatusNoContent)
	}

	return err
}
func getMessageByUser(c echo.Context) (err error) {
	userID := c.Param("id")

	if messages, count := dbGetMessageByUser(userID); count > 0 {
		n := MESS{messages}
		err = c.JSON(http.StatusOK, n)
	} else {
		err = c.NoContent(http.StatusNoContent)
	}

	return err
}

func getUserById(c echo.Context) (err error) {
	userId := c.Param("userId")

	if user, count := dbGetByUser(userId); count > 0 {
		n := USS{user}
		err = c.JSON(http.StatusOK, n)
	} else {
		err = c.NoContent(http.StatusNoContent)
	}
	return err
}

func getSalesOrders(c echo.Context) (err error) {

	if salesOrder, count := dbGetSalesOrders(); count > 0 {
		n := SOO{salesOrder}
		err = c.JSON(http.StatusOK, n)
	} else {
		err = c.NoContent(http.StatusNoContent)
	}
	return err
}
