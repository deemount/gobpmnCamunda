package elements

import (
	"github.com/deemount/gobpmnCamunda/pkg/camunda"
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
)

/*
 * Elementary
 */

// EndEvent ...
type EndEvent struct {
	camunda.CoreAttributes
	extension_elements.Extension_Elements
}

// TEndEvent ...
type TEndEvent struct {
	camunda.TCoreAttributes
	extension_elements.TExtension_Elements
}

// IntermediateCatchEvent ...
type IntermediateCatchEvent struct {
	camunda.CoreAttributes
	extension_elements.Extension_Elements
}

// TIntermediateCatchEvent ...
type TIntermediateCatchEvent struct {
	camunda.TCoreAttributes
	extension_elements.TExtension_Elements
}

// IntermediateThrowEvent ...
type IntermediateThrowEvent struct {
	extension_elements.Extension_Elements
}

// TIntermediateThrowEvent ...
type TIntermediateThrowEvent struct {
	extension_elements.TExtension_Elements
}

// StartEvent ...
type StartEvent struct {
	camunda.CoreAttributes
	extension_elements.Extension_Elements
	CamundaFormKey        string `xml:"camunda:formKey,attr,omitempty" json:"formKey,omitempty"`
	CamundaFormRef        string `xml:"camunda:formRef,attr,omitempty" json:"formRef,omitempty"`
	CamundaFormRefBind    string `xml:"camunda:formRefBinding,attr,omitempty" json:"formRefBind,omitempty"`
	CamundaFormRefVersion string `xml:"camunda:formRefVersion,attr,omitempty" json:"formRefVersion,omitempty"`
	CamundaInitiator      string `xml:"camunda:initiator,attr,omitempty" json:"init,omitempty"`
}

// TStartEvent ...
type TStartEvent struct {
	camunda.TCoreAttributes
	extension_elements.TExtension_Elements
	FormKey        string `xml:"formKey,attr,omitempty" json:"formKey,omitempty"`
	FormRef        string `xml:"formRef,attr,omitempty" json:"formRef,omitempty"`
	FormRefBind    string `xml:"formRefBinding,attr,omitempty" json:"formRefBind,omitempty"`
	FormRefVersion string `xml:"formRefVersion,attr,omitempty" json:"formRefVersion,omitempty"`
	Initiator      string `xml:"initiator,attr,omitempty" json:"init,omitempty"`
}
