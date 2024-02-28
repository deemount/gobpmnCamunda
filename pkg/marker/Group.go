package marker

import "github.com/deemount/gobpmnCamunda/pkg/extension_elements"

// NewGroup ...
func NewGroup() GroupRepository {
	return &Group{}
}

/*
 * @Setters
 */

/* Elements */

/** BPMN **/

// SetExtensionElements ...
func (group *Group) SetExtensionElements() {
	group.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (group Group) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &group.ExtensionElements[0]
}
