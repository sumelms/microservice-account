package http

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
	"github.com/sumelms/sumelms/user/pkg/adapter/database"
	storage "github.com/sumelms/sumelms/user/pkg/adapter/storage/gorm/user"
	"github.com/sumelms/sumelms/user/pkg/adapter/validator"
	"github.com/sumelms/sumelms/user/pkg/config"
	"github.com/sumelms/sumelms/user/pkg/context"
	"github.com/sumelms/sumelms/user/pkg/domain/user"
	handler "github.com/sumelms/sumelms/user/pkg/router/http/user"
)

type server struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewHttpServer() (*server, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		err = errors.Wrap(err, "Unable to load the configuration")
		return nil, err
	}

	db, err := database.Connect(cfg.Database)
	if err != nil {
		err = errors.Wrap(err, "Unable to connect with database")
		return nil, err
	}

	return &server{
		db:  db,
		cfg: cfg,
	}, nil
}

func (s server) Start() error {
	e := echo.New()

	e.Use(middleware.CORS())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		userRepository := storage.NewRepository(s.db)
		userService := user.NewService(userRepository)

		return func(c echo.Context) error {
			cc := &context.Context{
				Context: c,
				Service: userService,
			}
			return next(cc)
		}
	})

	// TODO Error handler

	e.Validator = validator.NewValidator()

	handler.NewHandler(e)

	if err := e.Start(s.cfg.Server.Http.Host); err != nil {
		err = errors.Wrap(err, "Unable to start the server")
		e.Logger.Error(err)
		return err
	}

	defer s.db.Close()

	return nil
}
