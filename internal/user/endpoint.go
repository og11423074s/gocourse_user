package user

import (
	"context"
	"github.com/og11423074s/go_lib_response/response"
)

type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	Endpoints struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}

	CreatReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	UpdateReq struct {
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
		Email     *string `json:"email"`
		Phone     *string `json:"phone"`
	}

	Config struct {
		LimPageDef string
	}
)

func makeCreateEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(CreatReq)

		// validations

		if req.FirstName == "" {
			return nil, response.BadRequest("first name is required")
		}

		if req.LastName == "" {
			return nil, response.BadRequest("last name is required")
		}

		// service
		user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email, req.Phone)
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}

		return response.Created("success", user, nil), nil
	}
}

/*
	func makeGetAllEndpoint(s Service, config Config) Controller {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			v := r.URL.Query()

			filters := Filters{
				FirstName: v.Get("first_name"),
				LastName:  v.Get("last_name"),
			}

			limit, _ := strconv.Atoi(v.Get("limit"))
			page, _ := strconv.Atoi(v.Get("page"))

			// select count(*) from users where first_name = ? and last_name = ?
			count, err := s.Count(filters)
			if err != nil {
				w.WriteHeader(500)
				json.NewEncoder(w).Encode(&Response{Status: 500, Error: err.Error()})
				return
			}

			metaResult, err := meta.New(page, limit, count, config.LimPageDef)

			if err != nil {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(&Response{Status: 400, Error: err.Error()})
				return
			}

			users, err := s.GetAll(filters, metaResult.Offset(), metaResult.Limit())
			if err != nil {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(&Response{Status: 400, Error: err.Error()})
				return
			}
			json.NewEncoder(w).Encode(&Response{Status: 200, Data: users, Meta: metaResult})
		}
	}

	func makeGetEndpoint(s Service) Controller {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			path := mux.Vars(r)
			id := path["id"]

			user, err := s.Get(id)

			if err != nil {
				w.WriteHeader(404)
				json.NewEncoder(w).Encode(&Response{Status: 404, Error: "user doesn't exist"})
				return
			}

			json.NewEncoder(w).Encode(&Response{Status: 200, Data: user})
		}
	}

	func makeUpdateEndpoint(s Service) Controller {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			var req UpdateReq

			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(&Response{Status: 400, Error: "invalid request format"})
				return
			}

			if req.FirstName != nil && *req.FirstName == "" {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(&Response{Status: 400, Error: "field name is required"})
				return
			}

			if req.LastName != nil && *req.LastName == "" {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(&Response{Status: 400, Error: "LastName name is required"})
				return
			}

			path := mux.Vars(r)
			id := path["id"]

			if err := s.Update(id, req.FirstName, req.LastName, req.Email, req.Phone); err != nil {
				w.WriteHeader(404)
				json.NewEncoder(w).Encode(&Response{Status: 400, Error: "user doesn't exist"})
				return
			}

			json.NewEncoder(w).Encode(&Response{Status: 200, Data: "user updated"})
		}
	}

	func makeDeleteEndpoint(s Service) Controller {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			path := mux.Vars(r)
			id := path["id"]

			if err := s.Delete(id); err != nil {
				w.WriteHeader(404)
				json.NewEncoder(w).Encode(&Response{Status: 404, Error: "user doesn't exist"})
				return
			}

			json.NewEncoder(w).Encode(&Response{Status: 200, Data: "user deleted"})

		}
	}
*/
func MakeEndpoints(s Service, config Config) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		/*Get:    makeGetEndpoint(s),
		GetAll: makeGetAllEndpoint(s, config),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),*/
	}
}
