package loop

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
)

// MultiInstanceLoopCharacteristics ...
type MultiInstanceLoopCharacteristics struct {
	CamundaAsyncBefore     bool                           `xml:"camunda:asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	CamundaAsyncAfter      bool                           `xml:"camunda:asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	CamundaCollection      string                         `xml:"camunda:collection,attr,omitempty" json:"collection,omitempty"`
	CamundaElementVariable string                         `xml:"camunda:elementVariable,attr,omitempty" json:"elementVariable,omitempty"`
	ExtensionElements      []attributes.ExtensionElements `xml:"bpmn:extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// TMultiInstanceLoopCharacteristics ...
type TMultiInstanceLoopCharacteristics struct {
	AsyncBefore       bool                            `xml:"asyncBefore,attr,omitempty" json:"asyncBefore,omitempty"`
	AsyncAfter        bool                            `xml:"asyncAfter,attr,omitempty" json:"asyncAfter,omitempty"`
	Collection        string                          `xml:"collection,attr,omitempty" json:"collection,omitempty"`
	ElementVariable   string                          `xml:"elementVariable,attr,omitempty" json:"elementVariable,omitempty"`
	ExtensionElements []attributes.TExtensionElements `xml:"extensionElements,omitempty" json:"extensionElements,omitempty"`
}

// ParticipantMultiplicity ...
type ParticipantMultiplicity struct{}

// TParticipantMultiplicity ...
type TParticipantMultiplicity struct{}

// StandardLoopCharacteristics ...
type StandardLoopCharacteristics struct{}

// TStandardLoopCharacteristics ...
type TStandardLoopCharacteristics struct{}
