package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gummy789j/bitgin-mock-server-staging/config"
	"github.com/gummy789j/bitgin-mock-server-staging/handler"
	"github.com/labstack/echo/v4"
)

var (
	key    string
	secret string
	port   string
)

func init() {

	flag.StringVar(&key, "k", "", "key")

	flag.StringVar(&secret, "s", "", "secret")

	flag.StringVar(&port, "p", "8080", "client serve port")

	flag.Usage = usage

	flag.Parse()

	if len(key) == 0 {
		panic("key is empty")
	}

	if len(secret) == 0 {
		panic("secret is empty")
	}

}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: bgsign [options] \n")
	fmt.Fprintf(os.Stderr, "  Currently, the following flags can be used\n")
	flag.PrintDefaults()
}

func main() {

	config.Initialize(key, secret)

	e := echo.New()

	e.POST("/v1/faas/pay", handler.FaasPayHandler)

	e.GET("/v1/faas/receipt", handler.FaasReceiptHandler)

	e.POST("/v1/mine/query", handler.MineQueryAddressesHandler)

	e.POST("/v1/mine/share", handler.MineShareHandler)

	e.Logger.Fatal(e.Start(":" + port))
}
