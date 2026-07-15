package main

import (
	"log"
	"net/http"
	"time"

	repo "github.com/Pranav2259/ecomm/internal/adapters/postgresql/sqlc"
	"github.com/Pranav2259/ecomm/internal/orders"
	"github.com/Pranav2259/ecomm/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

type app struct {
	config config
	db     *pgx.Conn
	//logger
	//client
}

type config struct {
	addr string
	db   dbconfig
}

type dbconfig struct {
	dsn string
	uri string
}

func (app *app) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All Good For Now"))
	})

	// first sample api
	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)
	r.Get("/products", productHandler.ListProducts)
	r.Get("/product/{id}", productHandler.FindProductByID)

	orderService := orders.NewService(repo.New(app.db), app.db)
	ordersHandler := orders.NewHandler(orderService)
	r.Post("/orders", ordersHandler.PlaceOrder)

	return r
}

func (app *app) run(h http.Handler) error {

	srv := http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		ReadTimeout:  (time.Second * 10),
		WriteTimeout: (time.Second * 30),
		IdleTimeout:  (time.Second * 60),
	}

	log.Printf("Server Started at %s", app.config.addr)

	return srv.ListenAndServe()

}
