package loop

type Timed struct {
	CurrentTimeInMS func() int
	ScheduleStep    func(f func(), delayInMS int)
}

func (t *Timed) Start(f func(), intervalInMS int) {
	r := &runner{
		currentTimeInMS: t.CurrentTimeInMS,
		scheduleStep:    t.ScheduleStep,
		f:               f,
		intervalInMS:    intervalInMS,
		lastStep:        t.CurrentTimeInMS() - intervalInMS,
	}
	r.step()
}

type runner struct {
	currentTimeInMS func() int
	scheduleStep    func(f func(), delay int)
	f               func()
	intervalInMS    int

	lastStep int
}

func (r *runner) step() {
	now := r.currentTimeInMS()
	delay := r.lastStep + 2*r.intervalInMS - now
	if delay < 0 {
		delay = 0
	}
	r.scheduleStep(r.step, delay)
	r.lastStep = now
	r.f()
}
