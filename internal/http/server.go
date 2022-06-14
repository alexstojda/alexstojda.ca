package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Server struct {
	SPAPath string
}

func NewServer(spaPath string) *Server {
	return &Server{
		SPAPath: spaPath,
	}
}

func (s *Server) StartServer() {
	router := gin.New()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Content-Type"},
		AllowMethods: []string{"POST", "DELETE"},
	}))

	// API ROUTES

	// SPA ROUTE
	log.Debug().Msg("Using " + s.SPAPath)
	router.Use(static.Serve("/", static.LocalFile(s.SPAPath, true)))

	err := router.Run()
	if err != nil {
		log.Error().Msgf("Web server startup failed with error %s", err)
	}
}
