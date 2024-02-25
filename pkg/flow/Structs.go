package flow

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
)

// Association ...
type Association struct {
	attributes.Attributes
}

// TAssociation ...
type TAssociation struct {
	attributes.TAttributes
}

// DataInputAssociation ...
type DataInputAssociation struct {
	attributes.Attributes
}

// TDataInputAssociation ...
type TDataInputAssociation struct {
	attributes.TAttributes
}

// MessageFlow ...
type MessageFlow struct {
	attributes.Attributes
}

// TMessageFlow ...
type TMessageFlow struct {
	attributes.TAttributes
}

// SequenceFlow ...
type SequenceFlow struct {
	attributes.Attributes
}

// TSequenceFlow ...
type TSequenceFlow struct {
	attributes.TAttributes
}
