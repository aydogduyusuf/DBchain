package api

import (
	"fmt"

	db "github.com/aydogduyusuf/DBchain/db/sqlc"
	"github.com/aydogduyusuf/DBchain/access_refresh_tokens"
	"github.com/aydogduyusuf/DBchain/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests
type Server struct {
	store db.Store
	router *gin.Engine
	tokenMaker access_refresh_tokens.Maker
	config util.Config
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config,store db.Store) (*Server, error) {
	tokenMaker, err := access_refresh_tokens.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store: store,
		tokenMaker: tokenMaker,
		config: config,
	}

	/* if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	} */

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))


	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes.POST("/users/tokens", server.deployToken)

	authRoutes.GET("/tokens/:id", server.getToken)
	authRoutes.GET("/tokens", server.listTokens)
	
	authRoutes.POST("/transactions", server.createTransaction)
	authRoutes.GET("/transactions", server.getTransaction)

	router.POST("/access_refresh_tokens/renew_access", server.renewAccessToken)
	
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}