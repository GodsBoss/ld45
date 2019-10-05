package loop

type Simple struct {
	ScheduleStep func(f func())
}

func (simple *Simple) Start(f func()) {
	sr := &simpleRunner{
		schedule: simple.ScheduleStep,
		f:        f,
	}
	sr.step()
}

type simpleRunner struct {
	schedule func(f func())
	f        func()
}

func (sr *simpleRunner) step() {
	sr.schedule(sr.step)
	sr.f()
}
