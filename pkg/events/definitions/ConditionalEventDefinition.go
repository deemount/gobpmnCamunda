package definitions

import impl "github.com/deemount/gobpmnTypes"

// NewConditionalEventDefinition ...
func NewConditionalEventDefinition() ConditionalEventDefinitionRepository {
	return &ConditionalEventDefinition{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamudnaVariableName ...
func (conditionalEventDefinition *ConditionalEventDefinition) SetCamundaVariableName(variableName string) {
	conditionalEventDefinition.CamundaVariableName = variableName
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaVariableName ...
func (conditionalEventDefinition ConditionalEventDefinition) GetCamundaVariableName() impl.STR_PTR {
	return &conditionalEventDefinition.CamundaVariableName
}
