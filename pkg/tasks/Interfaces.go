package tasks

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

/*
 * @Base
 */

// TaskCamundaBase ...
type TasksCamundaBase interface {
	SetCamundaAsyncBefore(asyncBefore bool)
	GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR
	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() gobpmnTypes.INT_PTR
}

// TasksBase ...
type TasksBase interface {
	TasksCamundaBase
	extension_elements.ExtensionElementsBaseElements
}

/*
 * @Repositories
 */

// BusinessRuleTaskRepository ...
type BusinessRuleTaskRepository interface {
	TasksBase
	SetCamundaClass(class string)
	GetCamundaClass() gobpmnTypes.STR_PTR
}

// ManualTaskRepository ...
type ManualTaskRepository interface {
	TasksBase
}

// ReceiveTaskRepository ...
type ReceiveTaskRepository interface {
	TasksBase
}

// ScriptTaskRepository ...
type ScriptTaskRepository interface {
	TasksBase
}

// SendTaskRepository ...
type SendTaskRepository interface {
	TasksBase
}

// ServiceTaskRepository ...
type ServiceTaskRepository interface {
	TasksBase
}

// TaskRepository ...
type TaskRepository interface {
	TasksBase
}

// UserTaskRepository ...
type UserTaskRepository interface {
	TasksBase
	SetCamundaFormKey(formKey string)
	GetCamundaFormKey() gobpmnTypes.STR_PTR
	SetCamundaAssignee(assignee string)
	GetCamundaAssignee() gobpmnTypes.STR_PTR
	SetCamundaCandidateUsers(users string)
	GetCamundaCandidateUsers() gobpmnTypes.STR_PTR
	SetCamundaCandidateGroups(groups string)
	GetCamundaCandidateGroups() gobpmnTypes.STR_PTR
	SetCamundaDueDate(due string)
	GetCamundaDueDate() gobpmnTypes.STR_PTR
	SetCamundaFollowUpDate(followUp string)
	GetCamundaFollowUpDate() gobpmnTypes.STR_PTR
	SetCamundaPriority(priority int)
	GetCamundaPriority() gobpmnTypes.INT_PTR
}
