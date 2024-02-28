package camunda

import gobpmnTypes "github.com/deemount/gobpmnTypes"

/*
 * @Base
 */

// CamundaBaseScriptElements ...
type CamundaBaseScriptElements interface {
	SetCamundaScript()
	GetCamundaScript() CAMUNDA_SCRIPT_PTR
}

// CamundaDefaultAttributes ...
type CamundaDefaultAttributes interface {
	SetCamundaAsyncBefore(asyncBefore bool)
	GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR
	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() gobpmnTypes.INT_PTR
}

type CamundaProcessAttributes interface {
	SetCamundaVersionTag(tag string)
	GetCamundaVersionTag() gobpmnTypes.STR_PTR
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() gobpmnTypes.INT_PTR
	SetCamundaTaskPriority(priority int)
	GetCamundaTaskPriority() gobpmnTypes.INT_PTR
	SetCamundaCandidateStarterGroups(groups string)
	GetCamundaCandidateStarterGroups() gobpmnTypes.STR_PTR
	SetCamundaCandiddateStarterUsers(users string)
	GetCamundaCandiddateStarterUsers() gobpmnTypes.STR_PTR
	SetCamundaHistoryTimeToLive(tolive string)
	GetCamundaHistoryTimeToLive() gobpmnTypes.STR_PTR
}

// ExtensionElementsRepository ...
type ExtensionElementsRepository interface {
	SetCamundaProperties()
	GetCamundaProperties() CAMUNDA_PROPERTIES_PTR
	SetCamundaFailedJobRetryCycle()
	GetCamundaFailedJobRetryCycle() CAMUNDA_FAILED_JOB_RETRY_CYCLE_PTR
	SetCamundaFormData()
	GetCamundaFormData() CAMUNDA_FORM_DATA_PTR
	SetCamundaInputOutput()
	GetCamundaInputOutput() CAMUNDA_IO_PTR
	SetCamundaTaskListener(num int)
	GetCamundaTaskListener(num int) CAMUNDA_TASK_LISTENER_PTR
	SetCamundaExecutionListener(num int)
	GetCamundaExecutionListener(num int) CAMUNDA_EXECUTION_LISTENER_PTR
	SetCamundaIn(num int)
	GetCamundaIn(num int) CAMUNDA_IN_PTR
	SetCamundaOut(num int)
	GetCamundaOut(num int) CAMUNDA_OUT_PTR
}

/*
 * @Repositories
 */

// CamundaConnectorRepository ...
type CamundaConnectorRepository interface{}

// CamundaConnectorIDRepository ...
type CamundaConnectorIDRepository interface {
	gobpmnTypes.IFBaseID
}

// CamundaConstraintRepository ...
type CamundaConstraintRepository interface {
	gobpmnTypes.IFBaseName
	gobpmnTypes.IFBaseConfig
}

// CamundaEntryRepository ...
type CamundaEntryRepository interface {
	gobpmnTypes.IFBaseValue
	gobpmnTypes.IFBaseKey
}

// CamundaExecutionListener ...
type CamundaExecutionListenerRepository interface {
	gobpmnTypes.IFBaseEvent
	gobpmnTypes.IFBaseClass
	CamundaBaseScriptElements
	SetDelegateExpression(expr string)
	GetDelegateExpression() gobpmnTypes.STR_PTR
	SetCamundaField(num int)
	GetCamundaField(num int) CAMUNDA_FIELD_PTR
}

// CamundaExpressionRepository ...
type CamundaExpressionRepository interface{}

// CamundaFailedJobyRetryCycleRepository  ...
type CamundaFailedJobRetryCycleRepository interface{}

// CamundaFieldRepository ...
type CamundaFieldRepository interface {
	gobpmnTypes.IFBaseName
	SetCamundaExpression()
	GetCamundaExpression() CAMUNDA_EXPRESSION_PTR
	SetCamundaString()
	GetCamundaString() CAMUNDA_STRING_PTR
}

// CamundaFormDataRepository ...
type CamundaFormDataRepository interface {
	SetCamundaFormField(num int)
	GetCamundaFormField(num int) CAMUNDA_FORM_FIELD_PTR
}

// CamundaFormField ...
type CamundaFormFieldRepository interface {
	gobpmnTypes.IFBaseID
	gobpmnTypes.IFBaseLabel
	gobpmnTypes.IFBaseType
	gobpmnTypes.IFBaseDefaultValue
	SetCamundaProperties()
	GetCamundaProperties() CAMUNDA_PROPERTIES_PTR
	SetCamundaValidation()
	GetCamundaValidation() CAMUNDA_VALIDATION_PTR
}

// CamundaInRepository ...
type CamundaInRepository interface{}

// CamundaInputOutputRepository ...
type CamundaInputOutputRepository interface {
	SetCamundaInputParameter(num int)
	GetCamundaInputParameter(num int) CAMUNDA_INPUT_PARAMETER_PTR
	SetCamundaOutputParameter(num int)
	GetCamundaOutputParameter(num int) CAMUNDA_OUTPUT_PARAMETER_PTR
}

// CamundaInputParameterRepository ...
type CamundaInputParameterRepository interface {
	CamundaBaseScriptElements
	gobpmnTypes.IFBaseLocalVariableName
	gobpmnTypes.IFBaseVariableAssignmentValue
	SetCamundaList()
	GetCamundaList() CAMUNDA_LIST_PTR
	SetCamundaMap()
	GetCamundaMap() CAMUNDA_MAP_PTR
}

// CamundaListRepository ...
type CamundaListRepository interface {
	SetCamundaValue(num int)
	GetCamundaValue(num int) CAMUNDA_VALUE_PTR
}

// CamundaMapRepository ...
type CamundaMapRepository interface {
	SetCamundaEntry(num int)
	GetCamundaEntry(num int) CAMUNDA_ENTRY_PTR
}

// CamundaOutRepository ...
type CamundaOutRepository interface{}

// CamundaOutputParameter ...
type CamundaOutputParameterRepository interface {
	CamundaBaseScriptElements
	gobpmnTypes.IFBaseLocalVariableName
	gobpmnTypes.IFBaseVariableAssignmentValue
	SetCamundaList()
	GetCamundaList() CAMUNDA_LIST_PTR
	SetCamundaMap()
	GetCamundaMap() CAMUNDA_MAP_PTR
}

// CamundaPropertiesRepository ...
type CamundaPropertiesRepository interface {
	SetCamundaProperty(num int)
	GetCamundaProperty(num int) CAMUNDA_PROPERTY_PTR
}

// CamundaPropertyRepository ...
type CamundaPropertyRepository interface {
	gobpmnTypes.IFBaseName
	gobpmnTypes.IFBaseValue
}

// CamundaScriptRepository ...
type CamundaScriptRepository interface {
	gobpmnTypes.IFBaseScriptFormat
}

// CamundaStringRepository ...
type CamundaStringRepository interface{}

// CamundaTaskListener ...
type CamundaTaskListenerRepository interface {
	gobpmnTypes.IFBaseEvent
	gobpmnTypes.IFBaseClass
	gobpmnTypes.IFBaseListenerID
	SetCamundaField(num int)
	GetCamundaField(num int) CAMUNDA_FIELD_PTR
}

// CamundaValidationRepository ...
type CamundaValidationRepository interface {
	SetCamundaConstraint(num int)
	GetCamundaConstraint(num int) CAMUNDA_CONSTRAINT_PTR
}

// CamundaValueRepository ...
type CamundaValueRepository interface {
	gobpmnTypes.IFBaseValue
}
