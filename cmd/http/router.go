package http

import (
	"fmt"

	"awesome-auth/configs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Config *configs.AppConfig
	DB     gorm.DB
}

func (s *Server) Run() {

	var mode string
	if s.Config.Mode == "production" || s.Config.Mode == "release" {
		mode = "release"
	} else {
		mode = "debug"
	}

	gin.SetMode(mode)
	engine := gin.New()

	v1Group := engine.Group("api/v1")
	s.DefineAuthRoutes(v1Group)

	err := engine.Run(fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port))
	if err != nil {
		panic(fmt.Sprintf("HTTP Server could not start: %s", err.Error()))
	}
}
