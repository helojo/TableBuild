package tool

import (
	"time"
)

type Clock struct {
	m_tick time.Time
}

func (clock *Clock) Start_tick() {
	clock.m_tick = time.Now()
}

// 返回秒
func (clock *Clock) End_tick() float64 {
	now_tick := time.Now()
	var dur_time time.Duration = now_tick.Sub(clock.m_tick)

	var elapsed_sec float64 = dur_time.Seconds()
	return elapsed_sec
}
