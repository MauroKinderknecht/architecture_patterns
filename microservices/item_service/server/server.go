package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/MauroKinderknecht/tech_talk/microservices/item_service/db"
	"github.com/MauroKinderknecht/tech_talk/microservices/item_service/handlers"
	"github.com/MauroKinderknecht/tech_talk/microservices/item_service/logger"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type server struct {
	log      *logger.Logger
	endpoint string

	backend *handlers.Backend
	router  *echo.Echo
	srv     *http.Server
}

func NewServer(b *handlers.Backend, db *db.Mongo, log *logger.Logger, port int) (*server, error) {
	router, err := configureRoutes(b, db, log)
	if err != nil {
		return nil, err
	}

	return &server{
		log:      log,
		endpoint: ":" + strconv.Itoa(port),
		backend:  b,
		router:   router,
	}, nil
}

func configureRoutes(backend *handlers.Backend, db *db.Mongo, logger *logger.Logger) (*echo.Echo, error) {
	logger.Info("setting up routes")

	swagger, err := handlers.GetSwagger()
	if err != nil {
		return nil, errors.Wrap(err, "error getting swagger")
	}

	router := echo.New()
	//router.Use(middleware.OapiRequestValidator(swagger))

	h := handlers.NewStrictHandler(backend, nil)
	basepath, _ := swagger.Servers.BasePath()
	handlers.RegisterHandlersWithBaseURL(router, h, basepath)

	return router, nil
}

func (s *server) Shutdown() error {
	return s.srv.Shutdown(context.Background())
}

// Start launches the web server and responds to requests
// it should be closed with Shutdown
func (s *server) Start() error {
	s.log.Info("Launching API server", logger.String("hostport", s.endpoint))

	s.srv = &http.Server{
		Addr:    s.endpoint,
		Handler: s.router,
	}

	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			s.log.Error("error in listen and serve", logger.Error(err))
		}
	}()
	return nil
}
