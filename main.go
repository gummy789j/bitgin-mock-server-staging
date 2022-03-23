package main

import (
	"os"

	"github.com/gummy789j/bitgin-mock-server-staging/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/v1/faas/pay", handler.FaasPayHandler)

	e.POST("/v1/faas/receipt", handler.FaasReceiptHandler)

	e.POST("/v1/mine/query", handler.MineQueryAddressesHandler)

	e.POST("/v1/mine/share", handler.MineShareHandler)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
