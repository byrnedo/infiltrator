package controllers

import (
	"github.com/byrnedo/apibase/controllers"
	. "github.com/byrnedo/apibase/logger"
	routes "github.com/byrnedo/apibase/routes"
	"github.com/byrnedo/apibase/utils"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"strconv"
)

// MainController handles incoming webhook payloads
type MainController struct {
	*controllers.BaseController
	restClient *utils.RestClient
}

// NewMainController utility method
func NewMainController() *MainController {
	return &MainController{
		BaseController: &controllers.BaseController{},
		restClient:     utils.NewRestClient(),
	}
}

// GetRoutes handles routing
func (pC *MainController) GetRoutes() []*routes.WebRoute {
	return []*routes.WebRoute{
		routes.NewWebRoute("HandlePortOpen", "/v1/connect", routes.GET, pC.CheckPortOpen),
		routes.NewWebRoute("HandleHttpStatus", "/v1/http", routes.GET, pC.CheckHTTPStatus),
	}
}

// CheckHTTPStatus checks the http status of the requested url
func (pC *MainController) CheckHTTPStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		url string
	)

	url = r.URL.Query().Get("url")

	if url == "" {
		w.WriteHeader(400)
		w.Write([]byte("url parameter missing"))
		return
	}

	if err := pC.restClient.Get(url); err != nil {
		w.WriteHeader(502)
		Error.Println(err)
		w.Write([]byte("could not connect to upstream"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(strconv.Itoa(pC.restClient.LastResponseStatus())))
}

// CheckPortOpen checks if the given host/port is open
func (pC *MainController) CheckPortOpen(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		port string
		host string
		err  error
	)

	host = r.URL.Query().Get("host")
	if host == "" {
		w.WriteHeader(400)
		w.Write([]byte("host paramter missing"))
		return
	}
	port = r.URL.Query().Get("port")
	if portInt, err := strconv.Atoi(port); err != nil || portInt < 1 {
		if err != nil {
			Error.Println(err)
		}
		w.WriteHeader(400)
		w.Write([]byte("invalid port"))
		return
	}

	conn, err := net.Dial("tcp", host+":"+port)

	if err != nil {
		w.WriteHeader(502)
		w.Write([]byte("could not connect to upstream"))
	} else {
		defer conn.Close()
		w.WriteHeader(200)
		w.Write([]byte("success"))
	}

}
