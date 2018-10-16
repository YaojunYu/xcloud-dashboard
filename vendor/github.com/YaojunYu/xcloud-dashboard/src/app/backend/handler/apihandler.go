package handler

import (
	"net/http"
	"github.com/emicklei/go-restful"
)

// APIHandler is a representation of API handler. Structure contains clientapi, Heapster clientapi and clientapi configuration.
type APIHandler struct {
	//iManager integration.IntegrationManager
	//cManager clientapi.ClientManager
	//sManager settings.SettingsManager
}
type Result struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func CreateHTTPAPIHandler() (http.Handler, error)  {
	apiHandler := APIHandler{}

	wsContainer := restful.NewContainer()
	wsContainer.EnableContentEncoding(true)

	apiV1Ws := new(restful.WebService)

	apiV1Ws.Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	wsContainer.Add(apiV1Ws)

	apiV1Ws.Route(
		apiV1Ws.POST("/login").
			To(apiHandler.handleLogin))

	return wsContainer, nil
}

// login api
func (apiHandler *APIHandler) handleLogin(request *restful.Request, response *restful.Response) {
	result := Result{Code: 200, Message: "Login succeed"}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}