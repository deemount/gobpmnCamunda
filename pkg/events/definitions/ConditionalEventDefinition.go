package definitions

import "github.com/deemount/gobpmnModels/pkg/impl"

// NewConditionalEventDefinition ...
func NewConditionalEventDefinition() ConditionalEventDefinitionRepository {
	return &ConditionalEventDefinition{}
}

/*
 * Default Setters
 */

/* Attributes */

/** Camunda **/

// SetCamudnaVariableName ...
func (conditionalEventDefinition *ConditionalEventDefinition) SetCamundaVariableName(variableName string) {
	conditionalEventDefinition.CamundaVariableName = variableName
}

/*
 * Default Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaVariableName ...
func (conditionalEventDefinition ConditionalEventDefinition) GetCamundaVariableName() impl.STR_PTR {
	return &conditionalEventDefinition.CamundaVariableName
}
