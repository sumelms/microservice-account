package json

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
	handler "github.com/sumelms/sumelms/user/pkg/router/json/user"
)

type router struct {
	db  *gorm.DB
	cfg *config.Config
}

func (r router) Start() error {
	e := echo.New()

	e.Use(middleware.CORS())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		userRepository := storage.NewRepository(r.db)
		userService := user.NewService(userRepository)

		return func(c echo.Context) error {
			cc := &context.Context{
				Context: c,
				Service: userService,
			}
			return next(cc)
		}
	})

	// TODO Error router

	e.Validator = validator.NewValidator()

	handler.NewHandler(e)

	if err := e.Start(r.cfg.GetPort()); err != nil {
		err = errors.Wrap(err, "Unable to start the server")
		e.Logger.Error(err)
		return err
	}

	defer r.db.Close()

	return nil
}

func NewRouter() (*router, error) {
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

	return &router{
		db:  db,
		cfg: cfg,
	}, nil
}
