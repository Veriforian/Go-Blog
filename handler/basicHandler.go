package handler

import (
	indexView "benjones/go-blog/view/index"

	"github.com/labstack/echo/v4"
)

type BasicHandler struct {
	Name string
}

func (handler BasicHandler) ShowHome(ctx echo.Context) error {
	name := handler.Name
	return render(ctx, indexView.ShowIndex(name))
}
