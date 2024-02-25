package data

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
)

/*
 * Default Setters
 */

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (dsr *DataStoreReference) SetExtensionElements() {
	dsr.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * Default Getters
 */

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (dsr DataStoreReference) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &dsr.ExtensionElements[0]
}
