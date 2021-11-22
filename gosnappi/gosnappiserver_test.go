package gosnappi_test

import (
	"fmt"

	"net/http"
	"testing"

	gosnappi "github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/open-traffic-generator/snappi/gosnappi/httpapi"

	"github.com/open-traffic-generator/snappi/gosnappi/httpapi/controllers"
	"github.com/open-traffic-generator/snappi/gosnappi/httpapi/interfaces"
)

//////////////////////////////////////////////////////////////////
// 			Strating gosnappi server with specific roure
//////////////////////////////////////////////////////////////////

func StartHttpServer() {
	fmt.Println("!!!!!!!!	Try to start gosnappi HTTP server	!!!!!!!!!!!!!")
	configHandler := NewConfigurationHandler()
	// Initiate other handler and append it
	controllers := []httpapi.HttpController{
		configHandler.GetController(),
	}
	router := httpapi.AppendRoutes(nil, controllers...)

	go func() {
		fmt.Println("Generated Http Server serving incoming HTTP requests on 127.0.0.1:50071.")
		if err := http.ListenAndServe("127.0.0.1:50071", router); err != nil {
			fmt.Println("Generated Http Server failed to serve incoming HTTP request.")
		}
	}()
}

//////////////////////////////////////////////////////////////////
// 		Initiating gosnappi handler and defined demo business logic
//////////////////////////////////////////////////////////////////

type configurationHandler struct {
	controller interfaces.ConfigurationController
}

func NewConfigurationHandler() interfaces.ConfigurationHandler {
	handler := new(configurationHandler)
	handler.controller = controllers.NewHttpConfigurationController(handler)
	return handler
}

func (h *configurationHandler) GetController() interfaces.ConfigurationController {
	return h.controller
}

func (h *configurationHandler) SetConfig(rbody gosnappi.Config, r *http.Request) gosnappi.SetConfigResponse {
	fmt.Println("SetConfig: We need to put our business logic ...")
	result := gosnappi.NewSetConfigResponse()
	result.StatusCode200()
	return result
}

func (h *configurationHandler) GetConfig(r *http.Request) gosnappi.GetConfigResponse {
	fmt.Println("GetConfig: We need to put our business logic ...")
	result := gosnappi.NewGetConfigResponse()
	result.StatusCode200()
	return result

}

//////////////////////////////////////////////////////////////////
// 		Test Code where init gosnappi server and access it through gosnappi package
//////////////////////////////////////////////////////////////////

func init() {
	StartHttpServer()
}

func config1(api gosnappi.GosnappiApi) gosnappi.Config {
	config := api.NewConfig()
	port1 := config.Ports().Add()
	port2 := config.Ports().Add()
	port1.SetName("port1")
	port1.SetLocation("location1")
	port2.SetName("port2")
	port2.SetLocation("location2")
	f1 := config.Flows().Add().SetName("f1")
	f2 := config.Flows().Add().SetName("f2")
	f1.TxRx().Port().SetTxName(port1.Name())
	f1.TxRx().Port().SetRxName(port2.Name())
	f1.Metrics().SetEnable(true)
	f2.TxRx().Port().SetTxName(port1.Name())
	f2.TxRx().Port().SetTxName(port2.Name())
	f2.Metrics().SetEnable(true)
	f1.Rate()
	return config
}

func TestSnappiServer(t *testing.T) {
	api := gosnappi.NewApi()
	api.NewHttpTransport().SetLocation("http://127.0.0.1:50071")
	config := config1(api)
	state, err := api.SetConfig(config)
	fmt.Println("state : ", state.ToJson())
	fmt.Println("err: ", err)
}
