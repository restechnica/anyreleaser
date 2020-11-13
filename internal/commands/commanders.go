package commands

import "os/exec"

type Commander interface {
	Output(name string, arg ...string) (output string, err error)
	Run(name string, arg ...string) (err error)
}

type ExecCommander struct{}

func NewExecCommander() *ExecCommander {
	return &ExecCommander{}
}

func (c ExecCommander) Output(name string, arg ...string) (string, error) {
	var output, err = exec.Command(name, arg...).Output()
	return string(output), err
}

func (c ExecCommander) Run(name string, arg ...string) error {
	return exec.Command(name, arg...).Run()
}
