package extension_elements

import (
	"github.com/deemount/gobpmnCamunda/pkg/camunda"
)

/*
 * @Base
 */

// AttributesElements ...
type AttributesBaseElements interface {
	SetExtensionElements()
	GetExtensionElements() EXTENSION_ELEMENTS_PTR
}

/*
 * @Repositories
 */

// ExtensionElementsRepository ...
type ExtensionElementsRepository interface {
	camunda.ExtensionElementsRepository
}
