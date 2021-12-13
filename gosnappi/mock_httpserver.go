package gosnappi

import (
	"io/ioutil"
	"log"
	"net/http"
)

type HttpServer struct {
	serverLocation string
	Api            GosnappiApi
	Config         Config
}

var (
	httpServer HttpServer = HttpServer{}
)

func StartMockHttpServer(location string) {
	httpServer.serverLocation = location
	httpServer.Api = NewApi()
	httpServer.Config = httpServer.Api.NewConfig()

	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			body, _ := ioutil.ReadAll(r.Body)
			httpServer.Config.FromJson(string(body))
			response := httpServer.Api.NewSetConfigResponse()
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			dataStr, _ := response.StatusCode200().ToJson()
			w.Write([]byte(dataStr))
		case http.MethodGet:
			// config := httpServer.Config
			response := httpServer.Api.NewGetConfigResponse()
			dataInp, _ := httpServer.Config.ToJson()
			response.StatusCode200().FromJson(dataInp)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			dataStr, _ := response.StatusCode200().ToJson()
			w.Write([]byte(dataStr))
		}
	})

	http.HandleFunc("/results/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodPost:
			body, _ := ioutil.ReadAll(r.Body)
			metricsReq := httpServer.Api.NewMetricsRequest()
			metricsReq.FromJson(string(body))
			if metricsReq.Choice() == MetricsRequestChoice.FLOW {
				flow_responses := httpServer.Api.NewGetMetricsResponse().StatusCode200()
				for _, flow_name := range metricsReq.Flow().FlowNames() {
					flow_rsp := flow_responses.FlowMetrics().Add()
					flow_rsp.SetName(flow_name)
					flow_rsp.SetBytesTx(1000)
					flow_rsp.SetBytesRx(1000)
				}
				response := httpServer.Api.NewGetMetricsResponse().StatusCode200()
				dataRsp, _ := flow_responses.ToJson()
				response.FromJson(dataRsp)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(dataRsp))
			} else if metricsReq.Choice() == MetricsRequestChoice.PORT {
				port_response := httpServer.Api.NewGetMetricsResponse().StatusCode200()
				for _, port_name := range metricsReq.Port().PortNames() {
					port_rsp := port_response.PortMetrics().Add()
					port_rsp.SetName(port_name)
					port_rsp.SetBytesTx(2000)
					port_rsp.SetBytesRx(2000)
				}
				response := httpServer.Api.NewGetMetricsResponse().StatusCode200()
				dataPort, _ := port_response.ToJson()
				response.FromJson(dataPort)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(dataPort))
			} else if metricsReq.Choice() == MetricsRequestChoice.BGPV4 {
				bgpv4_response := httpServer.Api.NewGetMetricsResponse().StatusCode200()
				for _, bgpv4_name := range metricsReq.Bgpv4().PeerNames() {
					bgpv4_rsp := bgpv4_response.Bgpv4Metrics().Add()
					bgpv4_rsp.SetName(bgpv4_name)
					bgpv4_rsp.SetRoutesAdvertised(80)
				}
				response := httpServer.Api.NewGetMetricsResponse().StatusCode200()
				dataBgp, _ := bgpv4_response.ToJson()
				response.FromJson(dataBgp)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(dataBgp))
			}
		}
	})

	go func() {
		if err := http.ListenAndServe(httpServer.serverLocation, nil); err != nil {
			log.Fatal("Server failed to serve incoming HTTP request.")
		}
	}()
}
