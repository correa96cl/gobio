package api

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
)

func (api *Api) BindRoutes() {

	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger, api.Sessions.LoadAndSave)

	csrfKey := []byte(os.Getenv("GOBID_CSRF_KEY"))

	// Instancia para TLS: cookies Secure + verificación estricta de Referer
	tlsCSRF := csrf.Protect(csrfKey, csrf.Secure(true), csrf.Path("/"))
	httpCSRF := csrf.Protect(csrfKey, csrf.Secure(false), csrf.Path("/"))

	// Middleware que elige la instancia según el protocolo real
	api.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Detecta TLS directo o a través de proxy inverso
			isTLS := r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https"

			if isTLS {
				tlsCSRF(next).ServeHTTP(w, r)
			} else {
				httpCSRF(next).ServeHTTP(w, csrf.PlaintextHTTPRequest(r))
			}
		})
	})
	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/csrftoken", api.HandleGetCSRFToken)
			r.Route("/users", func(r chi.Router) {
				r.Post("/signup", api.handleSignupUser)
				r.Post("/login", api.handleLoginUser)
				r.With(api.AuthMiddleware).Post("/logout", api.handleLogoutUser)
			})

			r.Route("/products", func(r chi.Router) {
				r.Post("/", api.handleCreateProduct)
			})
		})
	})
}
