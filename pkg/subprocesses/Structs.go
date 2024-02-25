package subprocesses

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	"github.com/deemount/gobpmnCamunda/pkg/camunda"
)

// Subprocesses ...
type Subprocesses struct{}

// TSubprocesses ...
type TSubprocesses struct{}

// AdHocSubProcess ...
type AdHocSubProcess struct {
	attributes.Attributes
	camunda.CoreAttributes
}

// TAdHocSubProcess ...
type TAdHocSubProcess struct {
	attributes.TAttributes
	camunda.TCoreAttributes
}

// CallActivity ...
type CallActivity struct {
	attributes.Attributes
	camunda.CoreAttributes
	CamundaCalledElementTenantID string `xml:"camunda:calledElementTenantId,attr,omitempty" json:"calledElementTenantId,omitempty"`
	CamundaVariableMappingClass  string `xml:"camunda:variableMappingClass,attr,omitempty" json:"variableMappingClass,omitempty"`
}

// TCallActivity ...
type TCallActivity struct {
	attributes.TAttributes
	camunda.TCoreAttributes
	CamundaCalledElementTenantID string `xml:"calledElementTenantId,attr,omitempty" json:"calledElementTenantId,omitempty"`
	CamundaVariableMappingClass  string `xml:"variableMappingClass,attr,omitempty" json:"variableMappingClass,omitempty"`
}

// SubProcess ...
type SubProcess struct {
	attributes.Attributes
	camunda.CoreAttributes
}

// TSubProcess ...
type TSubProcess struct {
	attributes.Attributes
	camunda.TCoreAttributes
}

// Transaction ...
type Transaction struct {
	attributes.Attributes
	camunda.CoreAttributes
}

// TTransaction ...
type TTransaction struct {
	attributes.TAttributes
	camunda.TCoreAttributes
}
