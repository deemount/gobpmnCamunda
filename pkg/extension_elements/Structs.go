package extension_elements

import (
	"github.com/deemount/gobpmnCamunda/pkg/camunda"
)

// Attributes ...
type Extension_Elements struct {
	ExtensionElements EXTENSION_ELEMENTS_SLC `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// TAttributes ...
type TExtension_Elements struct {
	ExtensionElements TEXTENSION_ELEMENTS_SLC `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// ExtensionElements ...
type ExtensionElements struct {
	camunda.ExtensionElements
}

// TExtensionElements ...
type TExtensionElements struct {
	camunda.ExtensionElements
}
