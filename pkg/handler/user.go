package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/og11423074s/go_lib_response/response"
	"github.com/og11423074s/gocourse_user/internal/user"
	"net/http"
	"strconv"
)

func NewUserHTTPServer(ctx context.Context, endpoints user.Endpoints) http.Handler {
	r := mux.NewRouter()

	opts := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Handle("/users", httptransport.NewServer(
		endpoint.Endpoint(endpoints.Create),
		decodeCreateUser, encodeResponse,
		opts...,
	)).Methods("POST")

	r.Handle("/users", httptransport.NewServer(
		endpoint.Endpoint(endpoints.GetAll),
		decodeGetAllUsers,
		encodeResponse,
		opts...,
	)).Methods("GET")

	r.Handle("/users/{id}", httptransport.NewServer(
		endpoint.Endpoint(endpoints.Get),
		decodeGetUser,
		encodeResponse,
		opts...,
	)).Methods("GET")

	r.Handle("/users/{id}", httptransport.NewServer(
		endpoint.Endpoint(endpoints.Update),
		decodeUpdateUser,
		encodeResponse,
		opts...,
	)).Methods("PATCH")

	r.Handle("/users/{id}", httptransport.NewServer(
		endpoint.Endpoint(endpoints.Delete),
		decodeDeleteUser,
		encodeResponse,
		opts...,
	)).Methods("DELETE")

	return r
}

func decodeUpdateUser(_ context.Context, r *http.Request) (interface{}, error) {
	var req user.UpdateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: %v", err.Error()))
	}

	p := mux.Vars(r)
	req.ID = p["id"]

	return req, nil
}

func decodeDeleteUser(_ context.Context, r *http.Request) (interface{}, error) {
	path := mux.Vars(r)
	req := user.DeleteReq{
		ID: path["id"],
	}

	return req, nil
}

func decodeGetUser(_ context.Context, r *http.Request) (interface{}, error) {

	p := mux.Vars(r)
	req := user.GetReq{
		ID: p["id"],
	}

	return req, nil
}

func decodeGetAllUsers(_ context.Context, r *http.Request) (interface{}, error) {

	v := r.URL.Query()

	limit, _ := strconv.Atoi(v.Get("limit"))
	page, _ := strconv.Atoi(v.Get("page"))

	req := user.GetAllReq{
		FirstName: v.Get("first_name"),
		LastName:  v.Get("last_name"),
		Limit:     limit,
		Page:      page,
	}

	return req, nil
}

func decodeCreateUser(_ context.Context, r *http.Request) (interface{}, error) {
	var req user.CreatReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: %v", err.Error()))
	}

	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	r := resp.(response.Response)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(r.StatusCode())
	return json.NewEncoder(w).Encode(r)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	r := err.(response.Response)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(r.StatusCode())
	_ = json.NewEncoder(w).Encode(r)
}
