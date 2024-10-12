package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type httpServer struct {
	server http.Server
}

func NewHttpServer() (*httpServer, *httprouter.Router) {
	router := httprouter.New()
	s := http.Server{
		Handler: router,
		//ReadTimeout:                  0,
		//WriteTimeout:                 0,
	}

	return &httpServer{s}, router
}

func (s *httpServer) Start(port int) error {
	s.server.Addr = fmt.Sprintf(":%d", port)

	log.Printf("Application is running on port %d\n", port)

	return s.server.ListenAndServe()
}
