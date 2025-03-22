package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/og11423074s/gocourse_user/pkg/bootstrap"
	"github.com/og11423074s/gocourse_user/pkg/handler"
	"os"

	"net/http"
	"time"

	"github.com/og11423074s/gocourse_user/internal/user"
)

func main() {

	// Load .env file
	_ = godotenv.Load()

	// Initialize logger
	logger := bootstrap.InitLogger()

	// Connect to database
	db, err := bootstrap.DBConnection()

	if err != nil {
		logger.Fatal(err)
	}

	pagLimDef := os.Getenv("PAGINATOR_LIMIT_DEFAULT")
	if pagLimDef == "" {
		logger.Fatal("PAGINATION_LIMIT_DEFAULT is not set")
	}

	ctx := context.Background()

	// User repository
	userRepo := user.NewRepo(logger, db)

	// User service
	userSrv := user.NewService(logger, userRepo)
	// User endpoints

	h := handler.NewUserHTTPServer(ctx, user.MakeEndpoints(userSrv, user.Config{LimPageDef: pagLimDef}))

	//router.HandleFunc("/users", userEnd.Create).Methods("POST")
	//router.HandleFunc("/users/{id}", userEnd.Get).Methods("GET")
	//router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	//router.HandleFunc("/users/{id}", userEnd.Update).Methods("PATCH")
	//router.HandleFunc("/users/{id}", userEnd.Delete).Methods("DELETE")

	port := os.Getenv("PORT")
	address := fmt.Sprintf("127.0.0.1:%s", port)

	srv := &http.Server{
		Handler:      accessControl(h),
		Addr:         address,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	errCh := make(chan error)
	go func() {
		logger.Println("listen in ", address)
		errCh <- srv.ListenAndServe()
	}()

	err = <-errCh
	if err != nil {
		logger.Fatal(err)
	}

}

func accessControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
