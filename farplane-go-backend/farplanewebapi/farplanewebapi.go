package farplanewebapi

import (
	"farplane/authentication"
	"farplane/database"
	"farplane/sessions"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func login(w http.ResponseWriter, req *http.Request) {
	authHeader := req.Header.Get("Authorization")
	fmt.Fprintf(w, "hello. %s that's it\n", authHeader)
	username, password, ok := req.BasicAuth()
	if ok {
		var hca authentication.HardcodedAuthenticator
		err := hca.ValidateUser(username, password)
		if err == nil {
			fmt.Fprintf(w, "Welcome\n")
			var fpSessions sessions.InMemFPSessions
			fpSessions.CreateSession(username)
			return
		}
	}
	fmt.Fprintf(w, "YOU ARE SUS AF\n")
}

func whoami(w http.ResponseWriter, req *http.Request) {
	authHeader := req.Header.Get("Authorization")
	fmt.Fprintf(w, "hello. %s that's it\n", authHeader)
	req.ParseForm()
	fmt.Fprintf(w, "testparam: URL %s QUERY %s\n", chi.URLParam(req, "testparam"), req.Form.Get("testparam"))
	var fpSessions sessions.InMemFPSessions
	fpSession, err := fpSessions.GetSession(authHeader)
	if err != nil {
		fmt.Fprint(w, "Could not obtain session for you\n")
	} else {
		fmt.Fprintf(w, "Welcome back, my Friend %s\n", fpSession.Username)
	}
	return
}

func signup(w http.ResponseWriter, req *http.Request) {

}
func CorsEnablerLocalhost(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
func StartServer(databaseI database.FarplaneDB) {
	r := chi.NewRouter()
	r.Use(CorsEnablerLocalhost)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Post("/login", login)
	r.Options("/login", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		fmt.Fprintf(w, "meh, that bout it")
	})
	r.Options("/whoami{testparam}", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "GET, that bout it")
	})
	r.Get("/whoami{testparam}", whoami)
	http.ListenAndServeTLS(":8090", "cert.pem", "key.pem", r)
}
