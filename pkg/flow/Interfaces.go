package flow

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
)

/*
 * @Repositories
 */

// AssociationRepository ...
type AssociationRepository interface {
	attributes.AttributesBaseElements
}

// DataInputAssociationRepository ...
type DataInputAssociationRepository interface {
	attributes.AttributesBaseElements
}

// MessageFlowRepository ...
type MessageFlowRepository interface {
	attributes.AttributesBaseElements
}

// SequenceFlowRepository ...
type SequenceFlowRepository interface {
	attributes.AttributesBaseElements
}
