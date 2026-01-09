package stats

import (
	"time"

	"github.com/Kai-qvq/perf-loadgen/internal/engine"
)

// Summary 是一次执行的统计汇总
type Summary struct {
	Total       int
	Successes   int
	Failures    int
	SuccessRate float64
	AvgLatency  time.Duration
}

// FromResults 将 engine 的结果聚合成 Summary
func FromResults(results []engine.Result) Summary {
	s := Summary{Total: len(results)}
	if s.Total == 0 {
		return s
	}

	var sum time.Duration
	for _, r := range results {
		sum += r.Latency
		if r.Success {
			s.Successes++
		} else {
			s.Failures++
		}
	}

	s.SuccessRate = float64(s.Successes) / float64(s.Total)
	s.AvgLatency = sum / time.Duration(s.Total)
	return s
}
