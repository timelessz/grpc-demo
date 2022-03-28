package calculate

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func NewHttpHandler(endpoint EndPointServer) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
	}
	r := mux.NewRouter()
	cal := r.Methods("GET").PathPrefix("/calculate").Subrouter()
	cal.Handle("/{type}/{a}/{b}", httptransport.NewServer(
		endpoint.CalculateEndPoint,
		decodeHTTPCalculateRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	r.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte(`{"status":"ok"}`))
	})
	return r
}

func decodeHTTPCalculateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	opType := vars["type"]
	a, _ := strconv.Atoi(vars["a"])
	b, _ := strconv.Atoi(vars["b"])
	return CalculateIn{
		A:    a,
		B:    b,
		Type: opType,
	}, nil
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	fmt.Println("errorEncoder", err.Error())
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errorWrap{Error: err.Error()})
}

type errorWrap struct {
	Error string `json:"error"`
}
