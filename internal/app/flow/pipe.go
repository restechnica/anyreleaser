package flow

import (
	"github.com/restechnica/anyreleaser/internal/app"
)

type Pipe interface {
	Run(ctx *app.Context) error
}
