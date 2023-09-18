package api

import (
	"main/database/db"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP request for our banking service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}

	router := gin.Default()
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAcount)
	router.GET("/accounts", server.listAcount)
	router.PATCH("/accounts/:id", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
