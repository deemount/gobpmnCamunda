package subprocesses

import (
	"github.com/deemount/gobpmnCamunda/pkg/camunda"
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

/*
 * @Base
 */

// SubprocessesBase ...
type SubprocessesBase interface {
	extension_elements.ExtensionElementsBaseElements
	camunda.CamundaDefaultAttributes
}

/*
 * @Repositories
 */

// AdHocSubProcessRepository ...
type AdHocSubProcessRepository interface {
	SubprocessesBase
}

// CallActivityRepository ...
type CallActivityRepository interface {
	SubprocessesBase

	SetCamundaCalledElementTenantID(tenantID string)
	GetCamundaCalledElementTenantID() gobpmnTypes.STR_PTR

	SetCamundaVariableMappingClass(class string)
	GetCamundaVariableMappingClass() gobpmnTypes.STR_PTR
}

// SubProcessRepository ...
type SubProcessRepository interface {
	SubprocessesBase
}

// TransactionRepository ...
type TransactionRepository interface {
	SubprocessesBase
}
