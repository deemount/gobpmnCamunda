package marker

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
)

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
	group.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * @Getters
 */

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (group Group) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &group.ExtensionElements[0]
}
