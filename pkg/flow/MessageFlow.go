package flow

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
)

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
	messageFlow.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * @Getters
 */

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetExtensionElements ...
func (messageFlow MessageFlow) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &messageFlow.ExtensionElements[0]
}
