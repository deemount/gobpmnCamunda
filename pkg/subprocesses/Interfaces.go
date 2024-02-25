package subprocesses

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	"github.com/deemount/gobpmnCamunda/pkg/camunda"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

/*
 * @Base
 */

// SubprocessesBase ...
type SubprocessesBase interface {
	attributes.AttributesBaseElements
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
	GetCamundaCalledElementTenantID() impl.STR_PTR

	SetCamundaVariableMappingClass(class string)
	GetCamundaVariableMappingClass() impl.STR_PTR
}

// SubProcessRepository ...
type SubProcessRepository interface {
	SubprocessesBase
}

// TransactionRepository ...
type TransactionRepository interface {
	SubprocessesBase
}
