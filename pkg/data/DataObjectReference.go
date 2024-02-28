package data

import "github.com/deemount/gobpmnCamunda/pkg/extension_elements"

/*
 * @Setters
 */

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (dor *DataObjectReference) SetExtensionElements() {
	dor.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (dor DataObjectReference) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &dor.ExtensionElements[0]
}
