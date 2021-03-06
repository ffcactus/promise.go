package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/enclosure/context"
)

// RefreshFan is the strategy to refresh fan.
type RefreshFan struct {
	name                string
	messageID           string
	description         string
	expectedExecutionMs uint64
}

// NewRefreshFan creates a new RefreshFan strategy.
func NewRefreshFan() *RefreshFan {
	return &RefreshFan{
		name:                "Refresh Fan",
		messageID:           "Promise.Enclosure.Action.Refresh.RefreshFan",
		description:         "Refresh enclosure fan components and their settings.",
		expectedExecutionMs: 5000,
	}
}

// Name returns the name of the strategy.
func (s *RefreshFan) Name() string {
	return s.name
}

// MessageID returns the message ID of the strategy.
func (s *RefreshFan) MessageID() string {
	return s.messageID
}

// Description returns the description of the strategy.
func (s *RefreshFan) Description() string {
	return s.description
}

// ExpectedExecutionMs returns the expected execution time in ms of the strategy.
func (s *RefreshFan) ExpectedExecutionMs() uint64 {
	return s.expectedExecutionMs
}

// Execute performs the operation of this strategy.
func (s *RefreshFan) Execute(c context.Refresh) {
	StepStart(c, s.name)
	slots, clientError := c.GetClient().FanSlot()
	if clientError != nil {
		// TODO we need process the alarm here.
		log.WithFields(log.Fields{
			"id": c.GetID(), "error": clientError,
		}).Warn("Strategy refresh fan failed, get fan slots failed.")
		StepError(c, s.name)
		return
	}
	enclosure, dbError := c.GetDB().RefreshFanSlot(c.GetID(), slots)
	if dbError != nil {
		log.WithFields(log.Fields{
			"id": c.GetID(), "error": clientError,
		}).Warn("Strategy refresh fan failed, DB refresh fan failed.")
	}
	c.UpdateEnclosure(enclosure)
	c.DispatchUpdateEvent()
	StepFinish(c, s.name)
	log.Info("Strategy refresh fan done.")
}
