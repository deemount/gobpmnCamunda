package flow

import "github.com/deemount/gobpmnCamunda/pkg/extension_elements"

// NewMessageFlow ...
func NewMessageFlow() MessageFlowRepository {
	return &MessageFlow{}
}

/*
 * @Setters
 */

/*** Make Elements ***/

/** BPMN **/

/*** Attributes ***/

// SetExtensionElements ...
func (messageFlow *MessageFlow) SetExtensionElements() {
	messageFlow.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetExtensionElements ...
func (messageFlow MessageFlow) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &messageFlow.ExtensionElements[0]
}
