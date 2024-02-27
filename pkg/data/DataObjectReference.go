package data

import "github.com/deemount/gobpmnCamunda/pkg/attributes"

/*
 * Default Setters
 */

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (dor *DataObjectReference) SetExtensionElements() {
	dor.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * Default Getters
 */

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (dor DataObjectReference) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &dor.ExtensionElements[0]
}
