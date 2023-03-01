package https

import (
	"fmt"
	"log"
	"net/http"

	"github.com/josephniel/josephniel.com/server/config"
	"github.com/josephniel/josephniel.com/server/router"
)

func NewRouter(conf config.Config) router.Router {
	return &httpRouter{
		config: conf,
	}
}

type httpRouter struct {
	config config.Config
}

func (r *httpRouter) Serve() error {
	log.Printf("Serving at %s:%v", r.config.Host, r.config.Port)

	mux := http.NewServeMux()

	r.registerRoutes(mux)

	return http.ListenAndServe(fmt.Sprintf("%s:%v", r.config.Host, r.config.Port), mux)
}

func (r *httpRouter) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("foo"))
	})
}
