package routers

import (
	"encoding/json"
	"fmt"
	"mama-chi/src/helper"
	"mama-chi/src/middleware"
	"mama-chi/src/models"
	"mama-chi/src/services"
	"net/http"

	"github.com/go-chi/render"

	"github.com/go-chi/chi"
)

// Index index
func Index() chi.Router {
	r := chi.NewRouter()
	r.Mount("/api/v1/auth", authRouter())
	return r
}

var pass = services.GetAllClient()

// A completely separate router for administrator routes
func authRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.BasicAuth("MyRealm", pass))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		helper.LogInfo("log Index")
		data := services.GetAllClient()
		render.JSON(w, r, data)
	})
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		username, password, _ := r.BasicAuth()
		fmt.Println(password)
		var users models.Users
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&users)
		if err != nil {
			render.Render(w, r, helper.ErrForbiden)
			return
		}
		fmt.Println(users)
		x, err := services.Login(users, username)
		if err != nil {
			render.Render(w, r, helper.ErrForbiden)
			return
		}
		json.NewEncoder(w).Encode(x)
	})
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("admin: view user id")))
	})
	return r
}
