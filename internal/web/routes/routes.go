package routes

import (
	"fmt"
	"forum/internal/service"
	"forum/internal/web/views"
	"net/http"
)

type Routes struct {
	service *service.Service
	view    *views.View
}

func NewRoutes(service *service.Service) (*Routes, error) {
	r := &Routes{
		service: service,
	}

	v, err := views.NewView()
	if err != nil {
		return nil, fmt.Errorf("Routes.NewRoutes: %w", err)
	}

	r.view = v
	return r, nil
}

func (r *Routes) InitRoutes(mux *http.ServeMux) {
	r.view.InitRoutes(mux)

	mux.HandleFunc("/", r.Home)

	mux.HandleFunc("/signup", r.SignUpPage)
	mux.HandleFunc("/register", r.SignUp)

	mux.HandleFunc("/signin", r.SignInPage)
	mux.HandleFunc("/login", r.SignIn)

	mux.HandleFunc("/logout", r.Logout)

	mux.HandleFunc("/tags", r.TagsPage)
	mux.HandleFunc("/tags/", r.Tags)

	mux.HandleFunc("/liked", r.UserLikedThreads)
	mux.HandleFunc("/mythreads", r.UserThreads)

	mux.HandleFunc("/thread", r.CreateThreadPage)
	mux.HandleFunc("/createthread", r.CreateThread)
	mux.HandleFunc("/thread/", r.ThreadPage)

	mux.HandleFunc("/comment", r.CreateComment)

	mux.HandleFunc("/like", r.Like)
}
