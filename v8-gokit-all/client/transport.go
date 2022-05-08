package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go-kit-demo/v8-gokit-all/auth"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

func MakeHttpHandler(endpoint endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/login").Name("login").Handler(kithttp.NewServer(
		endpoint,
		decodeLoginRequest,
		encodeLoginResponse,
	))
	r.Methods("GET").Path("/getUserInfo").Name("getUserInfo").Handler(kithttp.NewServer(
		endpoint,
		decodeLoginRequest,
		encodeLoginResponse,
	))
	return r
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	u, err := url.Parse(r.URL.RequestURI())
	if err != nil {
		return nil, err
	}
	fmt.Println(u.Path)
	switch u.Path {
	case "/login":
		username := r.URL.Query().Get("username")
		password := r.URL.Query().Get("password")
		return &auth.LoginRequest{
			Username: username,
			Password: password,
		}, nil
	case "/getUserInfo":
		userId := r.URL.Query().Get("user_id")
		userIdInt, err := strconv.Atoi(userId)
		if err != nil {
			return nil, err
		}
		return &auth.UserIdRequest{Id: int32(userIdInt)}, nil
	}
	return nil, nil
}

func encodeLoginResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	println("返回结果")
	println(reflect.TypeOf(response))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
