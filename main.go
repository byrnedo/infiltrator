package main

import (
	"fmt"
	"net/http"

	//"github.com/byrnedo/apibase/config"
	"flag"
	"github.com/byrnedo/apibase/controllers"
	. "github.com/byrnedo/apibase/logger"
	"github.com/byrnedo/apibase/middleware"
	webcontrollers "github.com/byrnedo/infiltrator/controllers"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "interface to bind to")
	flag.IntVar(&port, "port", 8080, "port to bind to")
	flag.Parse()
}

func main() {

	var externalRtr = httprouter.New()
	controllers.RegisterRoutes(externalRtr, webcontrollers.NewMainController())

	extHandlerChain := alice.New(
		middleware.LogTime,
		middleware.RecoverHandler,
		//middleware.AcceptJsonHandler,
	).Then(externalRtr)

	var listenAddr = fmt.Sprintf("%s:%d", host, port)
	Info.Printf("listening on " + listenAddr)
	if err := http.ListenAndServe(listenAddr, extHandlerChain); err != nil {
		panic("Failed to start server:" + err.Error())
	}
}
