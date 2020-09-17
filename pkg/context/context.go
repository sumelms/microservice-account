package context

import (
	"github.com/labstack/echo"
	"github.com/sumelms/sumelms/user/pkg/domain"
)

type Context struct {
	echo.Context
	Service domain.Service
}
