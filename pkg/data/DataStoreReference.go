package data

import "github.com/deemount/gobpmnCamunda/pkg/extension_elements"

/*
 * @Setters
 */

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (dsr *DataStoreReference) SetExtensionElements() {
	dsr.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (dsr DataStoreReference) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &dsr.ExtensionElements[0]
}
