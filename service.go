package automation

import (
	"fmt"

	"github.com/BlenderistDev/automation/interfaces"
)

// Service automation service
type Service struct {
	list map[string][]interfaces.Automation
}

// Start launch automation service
func (s *Service) Start(triggerChan chan interfaces.TriggerEvent, errChan chan error) {
	for {
		trigger := <-triggerChan
		err := s.handleTrigger(trigger)
		if err != nil {
			errChan <- err
		}
	}
}

// AddAutomation add automation to service
func (s *Service) AddAutomation(automation interfaces.Automation) {
	if len(s.list) == 0 {
		s.list = make(map[string][]interfaces.Automation)
	}
	for _, trigger := range automation.GetTriggers() {
		s.list[trigger] = append(s.list[trigger], automation)
	}
}

func (s *Service) handleTrigger(trigger interfaces.TriggerEvent) error {
	triggerName := trigger.GetName()
	automationList := s.list[triggerName]
	if automationList == nil {
		return fmt.Errorf("no automation for trigger %s", triggerName)
	}
	for _, automation := range automationList {
		fmt.Printf("TriggerEvent with type %s\n", triggerName)
		err := automation.Execute(trigger)
		if err != nil {
			return err
		}
	}
	return nil
}
