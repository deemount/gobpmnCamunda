package camunda

import "github.com/deemount/gobpmnModels/pkg/impl"

/*
 * @Base
 */

// CamundaBaseScriptElements ...
type CamundaBaseScriptElements interface {
	SetCamundaScript()
	GetCamundaScript() *CamundaScript
}

// CamundaDefaultAttributes ...
type CamundaDefaultAttributes interface {
	SetCamundaAsyncBefore(asyncBefore bool)
	GetCamundaAsyncBefore() impl.BOOL_PTR
	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() impl.BOOL_PTR
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() impl.INT_PTR
}

type CamundaProcessAttributes interface {
	SetCamundaVersionTag(tag string)
	GetCamundaVersionTag() impl.STR_PTR
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() impl.INT_PTR
	SetCamundaTaskPriority(priority int)
	GetCamundaTaskPriority() impl.INT_PTR
	SetCamundaCandidateStarterGroups(groups string)
	GetCamundaCandidateStarterGroups() impl.STR_PTR
	SetCamundaCandiddateStarterUsers(users string)
	GetCamundaCandiddateStarterUsers() impl.STR_PTR
	SetCamundaHistoryTimeToLive(tolive string)
	GetCamundaHistoryTimeToLive() impl.STR_PTR
}

/*
 * @Repositories
 */

// ExtensionElementsRepository ...
type ExtensionElementsRepository interface {
	SetCamundaProperties()
	GetCamundaProperties() *CamundaProperties
	SetCamundaFailedJobRetryCycle()
	GetCamundaFailedJobRetryCycle() *CamundaFailedJobRetryCycle
	SetCamundaFormData()
	GetCamundaFormData() *CamundaFormData
	SetCamundaInputOutput()
	GetCamundaInputOutput() *CamundaInputOutput
	SetCamundaTaskListener(num int)
	GetCamundaTaskListener(num int) *CamundaTaskListener
	SetCamundaExecutionListener(num int)
	GetCamundaExecutionListener(num int) *CamundaExecutionListener
	SetCamundaIn(num int)
	GetCamundaIn(num int) *CamundaIn
	SetCamundaOut(num int)
	GetCamundaOut(num int) *CamundaOut
}

// CamundaConnectorRepository ...
type CamundaConnectorRepository interface{}

// CamundaConnectorIDRepository ...
type CamundaConnectorIDRepository interface{}

// CamundaConstraintRepository ...
type CamundaConstraintRepository interface {
	impl.IFBaseName
	impl.IFBaseConfig
}

// CamundaEntryRepository ...
type CamundaEntryRepository interface {
	impl.IFBaseValue
	impl.IFBaseKey
}

// CamundaExecutionListener ...
type CamundaExecutionListenerRepository interface {
	impl.IFBaseEvent
	impl.IFBaseClass
	CamundaBaseScriptElements
	SetDelegateExpression(expr string)
	GetDelegateExpression() impl.STR_PTR
	SetCamundaField(num int)
	GetCamundaField(num int) *CamundaField
}

// CamundaExpressionRepository ...
type CamundaExpressionRepository interface{}

// CamundaFailedJobyRetryCycleRepository  ...
type CamundaFailedJobRetryCycleRepository interface{}

// CamundaFieldRepository ...
type CamundaFieldRepository interface {
	impl.IFBaseName
	SetCamundaExpression()
	GetCamundaExpression() *CamundaExpression
	SetCamundaString()
	GetCamundaString() *CamundaString
}

// CamundaFormDataRepository ...
type CamundaFormDataRepository interface {
	SetCamundaFormField(num int)
	GetCamundaFormField(num int) *CamundaFormField
}

// CamundaFormField ...
type CamundaFormFieldRepository interface {
	impl.IFBaseID
	impl.IFBaseLabel
	impl.IFBaseType
	impl.IFBaseDefaultValue
	SetCamundaProperties()
	GetCamundaProperties() *CamundaProperties
	SetCamundaValidation()
	GetCamundaValidation() *CamundaValidation
}

// CamundaInRepository ...
type CamundaInRepository interface{}

// CamundaInputOutputRepository ...
type CamundaInputOutputRepository interface {
	SetCamundaInputParameter(num int)
	GetCamundaInputParameter(num int) *CamundaInputParameter
	SetCamundaOutputParameter(num int)
	GetCamundaOutputParameter(num int) *CamundaOutputParameter
}

// CamundaInputParameterRepository ...
type CamundaInputParameterRepository interface {
	CamundaBaseScriptElements
	impl.IFBaseLocalVariableName
	impl.IFBaseVariableAssignmentValue
	SetCamundaList()
	GetCamundaList() *CamundaList
	SetCamundaMap()
	GetCamundaMap() *CamundaMap
}

// CamundaListRepository ...
type CamundaListRepository interface {
	SetCamundaValue(num int)
	GetCamundaValue(num int) *CamundaValue
}

// CamundaMapRepository ...
type CamundaMapRepository interface {
	SetCamundaEntry(num int)
	GetCamundaEntry(num int) *CamundaEntry
}

// CamundaOutRepository ...
type CamundaOutRepository interface{}

// CamundaOutputParameter ...
type CamundaOutputParameterRepository interface {
	CamundaBaseScriptElements
	impl.IFBaseLocalVariableName
	impl.IFBaseVariableAssignmentValue
	SetCamundaList()
	GetCamundaList() *CamundaList
	SetCamundaMap()
	GetCamundaMap() *CamundaMap
}

// CamundaPropertiesRepository ...
type CamundaPropertiesRepository interface {
	SetCamundaProperty(num int)
	GetCamundaProperty(num int) *CamundaProperty
}

// CamundaPropertyRepository ...
type CamundaPropertyRepository interface {
	impl.IFBaseName
	impl.IFBaseValue
}

// CamundaScriptRepository ...
type CamundaScriptRepository interface {
	impl.IFBaseScriptFormat
}

// CamundaStringRepository ...
type CamundaStringRepository interface{}

// CamundaTaskListener ...
type CamundaTaskListenerRepository interface {
	impl.IFBaseEvent
	impl.IFBaseClass
	impl.IFBaseListenerID
	SetCamundaField(num int)
	GetCamundaField(num int) *CamundaField
}

// CamundaValidationRepository ...
type CamundaValidationRepository interface {
	SetCamundaConstraint(num int)
	GetCamundaConstraint(num int) *CamundaConstraint
}

// CamundaValueRepository ...
type CamundaValueRepository interface {
	impl.IFBaseValue
}
