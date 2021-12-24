package api

import (
	"github.com/EgorSkurihin/Hokku/config"
	"github.com/EgorSkurihin/Hokku/store"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type APIServer struct {
	Echo         *echo.Echo
	addr         string
	logLevel     int
	store        store.Store
	sessionStore *sessions.CookieStore
}

func New(conf *config.Server, store store.Store) *APIServer {
	api := &APIServer{
		Echo:         echo.New(),
		addr:         conf.Addr,
		logLevel:     conf.LogLevel,
		sessionStore: sessions.NewCookieStore([]byte(conf.SessionKey)),
	}
	api.store = store
	return api
}

func (api *APIServer) Start() error {
	api.setupRoutes()
	return api.Echo.Start(api.addr)
}

func (api *APIServer) setupRoutes() {
	api.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	api.Echo.GET("/health", api.HealthCheck)
	api.Echo.GET("/hokkus", api.GetHokkus)
	api.Echo.GET("/hokkus/byTheme/:themeId", api.GetHokkusByTheme)
	api.Echo.GET("/hokkus/byAuthor/:authorId", api.GetHokkusByAuthor)
	api.Echo.GET("/hokku/:id", api.GetHokku)
	api.Echo.GET("/user/:id", api.GetUser)
	api.Echo.GET("/themes", api.GetThemes)
	api.Echo.POST("/user", api.PostUser)
	api.Echo.POST("/login", api.Login)

	restricted := api.Echo.Group("/restricted")
	//restricted.Use(session.Middleware(api.sessionStore))
	restricted.Use(api.authMiddleware)
	restricted.POST("/hokku", api.PostHokku)
	restricted.DELETE("/hokku/:id", api.DeleteHokku)
	restricted.PUT("/hokku/:id", api.PutHokku)
	restricted.DELETE("/user/:id", api.DeleteUser)
	restricted.PUT("/user/:id", api.PutUser)

	/* admin := api.Echo.Group("/admin")
	admin.POST("/theme", api.PostTheme)
	admin.DELETE("/theme/:id", api.DeleteTheme) */

	//swagger
	api.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
