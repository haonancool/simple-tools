package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

// Server simple http server
type Server struct {
	router     *mux.Router
	httpServer *http.Server
}

// NewServer create a server
func NewServer(addr string) *Server {
	router := mux.NewRouter()
	httpServer := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	server := &Server{
		router:     router,
		httpServer: httpServer,
	}
	return server
}

// Start start server
func (s *Server) Start() error {
	// init router
	s.router.HandleFunc("/getmyip", s.handleGetMyIP)

	// start http server
	err := s.httpServer.ListenAndServe()
	return err
}

func (s *Server) handleGetMyIP(w http.ResponseWriter, r *http.Request) {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad addr: %s", r.RemoteAddr)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(host))
}
