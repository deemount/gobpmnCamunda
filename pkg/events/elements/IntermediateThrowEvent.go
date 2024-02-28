package elements

import "github.com/deemount/gobpmnCamunda/pkg/extension_elements"

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
	intermediateThrowEvent.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (intermediateThrowEvent IntermediateThrowEvent) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &intermediateThrowEvent.ExtensionElements[0]
}
