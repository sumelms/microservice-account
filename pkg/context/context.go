package context

import (
	"github.com/labstack/echo"
	"github.com/sumelms/sumelms/user/pkg/domain/user"
)

type Context struct {
	echo.Context
	Service user.Service
}
