package subprocesses

import (
	"github.com/deemount/gobpmnCamunda/pkg/camunda"
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
)

// Subprocesses ...
type Subprocesses struct{}

// TSubprocesses ...
type TSubprocesses struct{}

/*
 * @Elementary
 */

// AdHocSubProcess ...
type AdHocSubProcess struct {
	extension_elements.Extension_Elements
	camunda.CoreAttributes
}

// TAdHocSubProcess ...
type TAdHocSubProcess struct {
	extension_elements.TExtension_Elements
	camunda.TCoreAttributes
}

// CallActivity ...
type CallActivity struct {
	extension_elements.Extension_Elements
	camunda.CoreAttributes
	CamundaCalledElementTenantID string `xml:"camunda:calledElementTenantId,attr,omitempty" json:"calledElementTenantId,omitempty"`
	CamundaVariableMappingClass  string `xml:"camunda:variableMappingClass,attr,omitempty" json:"variableMappingClass,omitempty"`
}

// TCallActivity ...
type TCallActivity struct {
	extension_elements.TExtension_Elements
	camunda.TCoreAttributes
	CamundaCalledElementTenantID string `xml:"calledElementTenantId,attr,omitempty" json:"calledElementTenantId,omitempty"`
	CamundaVariableMappingClass  string `xml:"variableMappingClass,attr,omitempty" json:"variableMappingClass,omitempty"`
}

// SubProcess ...
type SubProcess struct {
	extension_elements.Extension_Elements
	camunda.CoreAttributes
}

// TSubProcess ...
type TSubProcess struct {
	extension_elements.Extension_Elements
	camunda.TCoreAttributes
}

// Transaction ...
type Transaction struct {
	extension_elements.Extension_Elements
	camunda.CoreAttributes
}

// TTransaction ...
type TTransaction struct {
	extension_elements.TExtension_Elements
	camunda.TCoreAttributes
}
