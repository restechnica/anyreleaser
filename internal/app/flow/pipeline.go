package flow

import (
	"github.com/urfave/cli/v2"
)

type Pipeline struct {
	Pipes []Pipe
}

func (pipeline *Pipeline) Add(pipe Pipe) {
	pipeline.Pipes = append(pipeline.Pipes, pipe)
}

func (pipeline Pipeline) Run(ctx *cli.Context) (err error) {
	for _, pipe := range pipeline.Pipes {
		if err = pipe.Run(ctx); err != nil {
			return
		}
	}
	return
}
