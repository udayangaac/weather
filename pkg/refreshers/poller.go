package refreshers

import (
	"time"
)

type Poller interface {
	Poll()
	Stop()
}

type poller struct {
	timer       *time.Timer
	stopChan    chan struct{}
	triggerChan chan struct{}
	callback    func() (time.Duration, error)
}

func NewPoller(callback func() (time.Duration, error)) Poller {
	p := poller{
		triggerChan: make(chan struct{}),
		stopChan:    make(chan struct{}),
		callback:    callback,
	}
	p.init()
	return &p
}

func (p *poller) init() {
	go func(pInternal *poller) {
		for {
			select {
			case <-pInternal.stopChan:
				pInternal.timer.Stop()
				return
			case <-pInternal.triggerChan:
				duration, err := pInternal.callback()

				if err != nil {
					return
				}

				if duration == 0 {
					p.stopChan <- struct{}{}
				}

				pInternal.timer = time.AfterFunc(duration, func() {
					pInternal.triggerChan <- struct{}{}
				})
			}
		}
	}(p)
}

func (p *poller) Poll() {
	p.triggerChan <- struct{}{}
}

func (p *poller) Stop() {
	p.stopChan <- struct{}{}
}
