package flow

import "github.com/deemount/gobpmnCamunda/pkg/extension_elements"

/*
 * @Repositories
 */

// AssociationRepository ...
type AssociationRepository interface {
	extension_elements.ExtensionElementsBaseElements
}

// DataInputAssociationRepository ...
type DataInputAssociationRepository interface {
	extension_elements.ExtensionElementsBaseElements
}

// MessageFlowRepository ...
type MessageFlowRepository interface {
	extension_elements.ExtensionElementsBaseElements
}

// SequenceFlowRepository ...
type SequenceFlowRepository interface {
	extension_elements.ExtensionElementsBaseElements
}
