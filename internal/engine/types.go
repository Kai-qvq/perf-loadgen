package engine

import "time"

type Result struct {
	Latency time.Duration
	Success bool
}

type Runner interface {
	Run() ([]Result, error)
}
