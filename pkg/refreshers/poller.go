package refreshers

import (
	"time"
)

// Poller is an interface for polling and stopping polling.
type Poller interface {
	Poll()
	Stop()
}

// poller is a struct that implements the Poller interface.
type poller struct {
	timer       *time.Timer
	stopChan    chan struct{}
	triggerChan chan struct{}
	callback    func() (time.Duration, error)
}

// NewPoller creates a new Poller instance with the provided callback function.
func NewPoller(callback func() (time.Duration, error)) Poller {
	p := poller{
		triggerChan: make(chan struct{}),
		stopChan:    make(chan struct{}),
		callback:    callback,
	}
	p.init()
	return &p
}

// init initializes the polling mechanism in a goroutine.
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

// Poll triggers the polling mechanism.
func (p *poller) Poll() {
	p.triggerChan <- struct{}{}
}

// Stop stops the polling mechanism.
func (p *poller) Stop() {
	p.stopChan <- struct{}{}
}
