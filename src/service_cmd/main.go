package main

import (
	"github.com/bladedancer/ratelimit/src/service_cmd/runner"
	"github.com/bladedancer/ratelimit/src/settings"
)

func main() {
	runner := runner.NewRunner(settings.NewSettings())
	runner.Run()
}
