package main

import (
	"fmt"
	"os"

	"github.com/qwenode/convergen/pkg/config"
	"github.com/qwenode/convergen/pkg/runner"
)

//go:generate go install .
func main() {
	var conf config.Config
	if err := conf.ParseArgs(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := runner.Run(conf); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
