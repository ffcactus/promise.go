package dto

import (
	"promise/base"
	"promise/task/object/errorResp"
	"promise/task/object/model"
)

// PostTaskRequest Post task request DTO.
type PostTaskRequest struct {
	MessageID     string                `json:"MessageID,omitempty"`
	Name          string                `json:"Name"`
	Description   string                `json:"Description,omitempty"`
	CreatedByName string                `json:"CreatedByName"`
	CreatedByURI  string                `json:"CreatedByURI"`
	TargetName    string                `json:"TargetName"`
	TargetURI     string                `json:"TargetURI"`
	TaskSteps     []PostTaskStepRequest `json:"TaskSteps"`
}

// NewInstance creates a new instance.
func (PostTaskRequest) NewInstance() base.RequestInterface {
	return new(PostTaskRequest)
}

// IsValid return if the request is valid.
func (dto *PostTaskRequest) IsValid() *base.ErrorResponse {
	if len(dto.TaskSteps) == 0 {
		return errorResp.NewErrorResponseTaskNoStep()
	}
	return nil
}

// String return the name for debug.
func (dto PostTaskRequest) String() string {
	return dto.Name
}

// ToModel convert the DTO to model.
func (dto PostTaskRequest) ToModel() base.ModelInterface {
	m := model.Task{}
	m.Category = base.CategoryTask
	m.MessageID = dto.MessageID
	m.Name = dto.Name
	m.Description = dto.Description
	m.ExecutionState = model.ExecutionStateReady
	m.CreatedByName = dto.CreatedByName
	m.CreatedByURI = dto.CreatedByURI
	m.TargetName = dto.TargetName
	m.TargetURI = dto.TargetURI
	m.CurrentStep = dto.TaskSteps[0].Name
	m.Percentage = 0
	m.ExecutionResult.State = model.ExecutionResultStateUnknown
	m.ExpectedExecutionMs = 0
	for i := range dto.TaskSteps {
		m.TaskSteps = append(m.TaskSteps, *dto.TaskSteps[i].ToModel())
		// The task execution time equals to the sum of every steps'.
		m.ExpectedExecutionMs += dto.TaskSteps[i].ExpectedExecutionMs
	}
	return &m
}
