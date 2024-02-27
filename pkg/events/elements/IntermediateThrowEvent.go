package elements

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
)

// NewIntermediateThrowEvent ...
func NewIntermediateThrowEvent() IntermediateThrowEventRepository {
	return &IntermediateThrowEvent{}
}

/*
 * @Setters
 */

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (intermediateThrowEvent *IntermediateThrowEvent) SetExtensionElements() {
	intermediateThrowEvent.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * @Getters
 */

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (intermediateThrowEvent IntermediateThrowEvent) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &intermediateThrowEvent.ExtensionElements[0]
}
