package definitions

import impl "github.com/deemount/gobpmnTypes"

/*
 * @Repositories
 */

// CancelEventDefinitionRepository ...
type CancelEventDefinitionRepository interface{}

// CompensateEventDefinitionRepository ...
type CompensateEventDefinitionRepository interface{}

// ConditionalEventDefinitionRepository ...
type ConditionalEventDefinitionRepository interface {
	SetCamundaVariableName(variableName string)
	GetCamundaVariableName() impl.STR_PTR
}

// ErrorEventDefinitionRepository ...
type ErrorEventDefinitionRepository interface{}

// EscalationEventDefinitionRepository ...
type EscalationEventDefinitionRepository interface{}

// LinkEventDefinitionRepository ...
type LinkEventDefinitionRepository interface{}

// MessageEventDefinitionRepository ...
type MessageEventDefinitionRepository interface{}

// SignalEventDefinitionRepository ...
type SignalEventDefinitionRepository interface{}

// TerminateEventDefinitionRepository ...
type TerminateEventDefinitionRepository interface{}

// TimerEventDefinitionRepository ...
type TimerEventDefinitionRepository interface{}
