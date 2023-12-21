package server

import (
	"backend/internal/backend/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	route  *gin.Engine
	config *config.Config
}
