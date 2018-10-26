package routes

import (
	"github.com/Muart-C/webcbo/controllers"
	admin "github.com/Muart-C/webcbo/controllers"
	mw "github.com/Muart-C/webcbo/routes/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gobuffalo/envy"
	"github.com/gorilla/csrf"
	"log"
	"net/http"
	"time"
)

// New creates a new router instance and server
func ServeRouter() error {
	mux := HandleRoutes()

	addr := ":" + envy.Get("PORT", "9001")
	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// production server will use nginx + letsencrypt
	log.Println("Server is running")
	return server.ListenAndServe()
}

func HandleRoutes() *chi.Mux {
	r := chi.NewRouter()
	env := envy.Get("ENVIRONMENT", "development")

	// Middlewares
	r.Use(mw.MethodOverride)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.RedirectSlashes)
	r.Use(mw.Recoverer)
	r.Use(middleware.DefaultCompress)

	// Base
	r.NotFound(controllers.NotFoundHandler)
	r.Get("/public/*", func(w http.ResponseWriter, r *http.Request) {
		if env == "production" {
			w.Header().Set("Vary", "Accept-Encoding")
			w.Header().Set("Cache-Control", "public, max-age=604800")
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		fs := http.StripPrefix("/public", http.FileServer(http.Dir("public")))

		fs.ServeHTTP(w, r)
	})
	serveSingle(r, "/favicon.ico", "public/favicon.ico")

	r.Route("/", func(r chi.Router) {
		r.NotFound(controllers.NotFoundHandler)

		csrfKey := []byte(envy.Get("CSRF_KEY", "g4k827b582367a77cb27d1e5dc268912"))

		if env != "test" {
			r.Use(csrf.Protect(
				csrfKey,
				csrf.Secure(false), // change to true after switching to https
			))
		}

		// home signup
		//r.Get("/", h(controllers.Home))

		// Admin auth
		//r.Get("/admin/sessions/new", h(admin.GetLogin))
		//r.Post("/admin/sessions", h(admin.PostLogin))
		//r.Delete("/admin/sessions", h(admin.Logout))

		r.Route("/admin", func(r chi.Router) {
			//r.Use(mw.RequireAdminAuth)

			// Dashboard
			r.Get("/", h(admin.Dashboard))

			// Support Messages
			//r.Route("/support-messages", func(r chi.Router) {
			//      r.Get("/", h(admin.ListMessages))
			//      r.Delete("/{ID}", h(admin.DeleteMessage))
			//})

			// Projects
			r.Route("/users", func(r chi.Router) {
				r.Get("/", h(admin.ListUsers))
				//r.Get("/new", h(admin.NewProject))
				//r.Post("/", h(admin.CreateProject))
				//
				//r.Route("/{ID}", func(r chi.Router) {
				//      r.Use(admin.PageContext)
				//      r.Get("/", h(admin.GetPproject))
				//      r.Get("/edit", h(admin.EditProject))
				//      r.Put("/", h(admin.UpdateProject))
				//      //r.Delete("/", h(admin.DeleteProject))
			})
		})

	})

	//})

	return r
}

func h(fn controllers.Handler) http.HandlerFunc {
	return controllers.Handler(fn).ServeHTTP
}

func serveSingle(mux *chi.Mux, pattern, filename string) {
	mux.Get(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}
