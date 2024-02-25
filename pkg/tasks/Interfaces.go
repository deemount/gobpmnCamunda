package tasks

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

/*
 * @Base
 */

// TaskCamundaBase ...
type TasksCamundaBase interface {
	SetCamundaAsyncBefore(asyncBefore bool)
	GetCamundaAsyncBefore() impl.BOOL_PTR
	SetCamundaAsyncAfter(asyncAfter bool)
	GetCamundaAsyncAfter() impl.BOOL_PTR
	SetCamundaJobPriority(priority int)
	GetCamundaJobPriority() impl.INT_PTR
}

// TasksBaseCoreElements ...
type TasksBaseCoreElements interface {
	SetExtensionElements()
	GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR
}

// TasksBase ...
type TasksBase interface {
	TasksCamundaBase
	TasksBaseCoreElements
}

/*
 * @Repositories
 */

// BusinessRuleTaskRepository ...
type BusinessRuleTaskRepository interface {
	TasksBase
	SetCamundaClass(class string)
	GetCamundaClass() impl.STR_PTR
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
	GetCamundaFormKey() impl.STR_PTR
	SetCamundaAssignee(assignee string)
	GetCamundaAssignee() impl.STR_PTR
	SetCamundaCandidateUsers(users string)
	GetCamundaCandidateUsers() impl.STR_PTR
	SetCamundaCandidateGroups(groups string)
	GetCamundaCandidateGroups() impl.STR_PTR
	SetCamundaDueDate(due string)
	GetCamundaDueDate() impl.STR_PTR
	SetCamundaFollowUpDate(followUp string)
	GetCamundaFollowUpDate() impl.STR_PTR
	SetCamundaPriority(priority int)
	GetCamundaPriority() impl.INT_PTR
}
