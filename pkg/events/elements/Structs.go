package elements

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	"github.com/deemount/gobpmnCamunda/pkg/camunda"
)

/*
 * Elementary
 */

// EndEvent ...
type EndEvent struct {
	camunda.CoreAttributes
	attributes.Attributes
}

// TEndEvent ...
type TEndEvent struct {
	camunda.TCoreAttributes
	attributes.TAttributes
}

// IntermediateCatchEvent ...
type IntermediateCatchEvent struct {
	camunda.CoreAttributes
	attributes.Attributes
}

// TIntermediateCatchEvent ...
type TIntermediateCatchEvent struct {
	camunda.TCoreAttributes
	attributes.TAttributes
}

// IntermediateThrowEvent ...
type IntermediateThrowEvent struct {
	attributes.Attributes
}

// TIntermediateThrowEvent ...
type TIntermediateThrowEvent struct {
	attributes.TAttributes
}

// StartEvent ...
type StartEvent struct {
	camunda.CoreAttributes
	attributes.Attributes
	CamundaFormKey        string `xml:"camunda:formKey,attr,omitempty" json:"formKey,omitempty"`
	CamundaFormRef        string `xml:"camunda:formRef,attr,omitempty" json:"formRef,omitempty"`
	CamundaFormRefBind    string `xml:"camunda:formRefBinding,attr,omitempty" json:"formRefBind,omitempty"`
	CamundaFormRefVersion string `xml:"camunda:formRefVersion,attr,omitempty" json:"formRefVersion,omitempty"`
	CamundaInitiator      string `xml:"camunda:initiator,attr,omitempty" json:"init,omitempty"`
}

// TStartEvent ...
type TStartEvent struct {
	camunda.TCoreAttributes
	attributes.TAttributes
	FormKey        string `xml:"formKey,attr,omitempty" json:"formKey,omitempty"`
	FormRef        string `xml:"formRef,attr,omitempty" json:"formRef,omitempty"`
	FormRefBind    string `xml:"formRefBinding,attr,omitempty" json:"formRefBind,omitempty"`
	FormRefVersion string `xml:"formRefVersion,attr,omitempty" json:"formRefVersion,omitempty"`
	Initiator      string `xml:"initiator,attr,omitempty" json:"init,omitempty"`
}
