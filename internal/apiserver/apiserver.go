package apiserver

import (
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	config *Config
	router *gin.Engine
}

func NewAPIServer(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: gin.New(),
	}
}

func(s *APIServer) Start() {
	s.configureRouter()

	s.router.Run(s.config.BindAddress)
}

func(s *APIServer) configureRouter() {
	s.router.Use(gin.Logger())
	s.router.Use(headerCheck())
	s.router.GET("/when/:year", getNumbersOfDays)

}

