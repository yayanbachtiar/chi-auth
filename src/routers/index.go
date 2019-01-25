package routers

import (
	"fmt"
	"mama-chi/src/middleware"
	"mama-chi/src/services"
	"net/http"

	"github.com/go-chi/render"

	"github.com/go-chi/chi"
)

// Index index
func Index() chi.Router {
	r := chi.NewRouter()
	r.Mount("/api/v1", adminRouter())
	return r
}

var pass = services.GetAllClient()

// A completely separate router for administrator routes
func adminRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.New("MyRealm", pass))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data := services.GetAllClient()
		render.JSON(w, r, data)
	})
	r.Get("/accounts", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin: list accounts.."))
	})
	r.Get("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("admin: view user id %v", chi.URLParam(r, "userId"))))
	})
	return r
}
