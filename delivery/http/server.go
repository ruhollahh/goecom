package httpserver

import (
	"context"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/ruhollahh/go-ecom/delivery/http/handler/admin/admin"
	checkhandler "github.com/ruhollahh/go-ecom/delivery/http/handler/check"
	producthandler "github.com/ruhollahh/go-ecom/delivery/http/handler/product"
	userhandler "github.com/ruhollahh/go-ecom/delivery/http/handler/user/user"
	"github.com/ruhollahh/go-ecom/delivery/http/mid"
	"github.com/ruhollahh/go-ecom/internal/service"
	"github.com/ruhollahh/go-ecom/pkg/logger"
	session "github.com/spazzymoto/echo-scs-session"
	echoSwagger "github.com/swaggo/echo-swagger"
	"time"
)

type Config struct {
	Build   string
	APIPort string
	APIHost string
}

type Server struct {
	cfg             Config
	router          *echo.Echo
	shutdownTimeout time.Duration
	log             *logger.Logger
	sessionManager  *scs.SessionManager
	checkHandler    *checkhandler.Handler
	adminHandler    *adminhandler.Handler
	userHandler     *userhandler.Handler
	productHandler  *producthandler.Handler
}

func New(cfg Config, log *logger.Logger, db *sqlx.DB, services service.Service, sessionManager *scs.SessionManager) *Server {
	return &Server{
		cfg:            cfg,
		router:         echo.New(),
		log:            log,
		sessionManager: sessionManager,
		checkHandler:   checkhandler.New(cfg.Build, db),
		adminHandler:   adminhandler.New(services.AdminSvc),
		userHandler:    userhandler.New(services.UserSvc),
		productHandler: producthandler.New(services.ProductSvc),
	}
}

func (h *Server) Serve() error {
	h.registerRoutes()

	return h.router.Start(fmt.Sprintf("%s:%s", h.cfg.APIHost, h.cfg.APIPort))
}

func (h *Server) registerRoutes() {
	v1 := h.router.Group("/v1")
	v1.Use(mid.Tracer())
	v1.Use(mid.Logger(h.log))
	v1.Use(mid.Errors(h.log))
	v1.Use(session.LoadAndSave(h.sessionManager))
	v1.Use(mid.Panics())
	v1.GET("/swagger/*", echoSwagger.WrapHandler)

	h.checkHandler.RegisterRoutes(v1)

	h.adminHandler.RegisterRoutes(v1)

	h.userHandler.RegisterRoutes(v1)

	h.productHandler.RegisterRoutes(v1)
}

func (h *Server) Shutdown(ctx context.Context) error {
	return h.router.Shutdown(ctx)
}

func (h *Server) Close() error {
	return h.router.Close()
}
