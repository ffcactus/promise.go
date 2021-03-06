package strategy

import (
	log "github.com/sirupsen/logrus"
	"promise/enclosure/context"
)

// RefreshManager is the strategy to refresh manager.
type RefreshManager struct {
	name                string
	messageID           string
	description         string
	expectedExecutionMs uint64
}

// NewRefreshManager creates a new RefreshManager strategy.
func NewRefreshManager() *RefreshManager {
	return &RefreshManager{
		name:                "Refresh Manager",
		messageID:           "Promise.Enclosure.Action.Refresh.RefreshManager",
		description:         "Refresh enclosure manager components and their settings.",
		expectedExecutionMs: 5000,
	}
}

// Name returns the name of the strategy.
func (s *RefreshManager) Name() string {
	return s.name
}

// MessageID returns the message ID of the strategy.
func (s *RefreshManager) MessageID() string {
	return s.messageID
}

// Description returns the description of the strategy.
func (s *RefreshManager) Description() string {
	return s.description
}

// ExpectedExecutionMs returns the expected execution time in ms of the strategy.
func (s *RefreshManager) ExpectedExecutionMs() uint64 {
	return s.expectedExecutionMs
}

// Execute performs the operation of this strategy.
func (s *RefreshManager) Execute(c context.Refresh) {
	StepStart(c, s.name)
	slots, clientError := c.GetClient().ManagerSlot()
	if clientError != nil {
		// TODO we need process the alarm here.
		log.WithFields(log.Fields{
			"id": c.GetID(), "error": clientError,
		}).Warn("Strategy refresh manager failed, get manager slots failed.")
		StepError(c, s.name)
		return
	}
	enclosure, dbError := c.GetDB().RefreshManagerSlot(c.GetID(), slots)
	if dbError != nil {
		StepError(c, s.name)
		log.WithFields(log.Fields{
			"id": c.GetID(), "error": clientError,
		}).Warn("Strategy refresh manager failed, DB refresh manager failed.")
	}
	c.UpdateEnclosure(enclosure)
	c.DispatchUpdateEvent()
	StepFinish(c, s.name)
	log.Info("Strategy refresh manager done.")
}
