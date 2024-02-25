package tasks

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	"github.com/deemount/gobpmnModels/pkg/impl"
)

// NewUserTask ...
func NewUserTask() UserTaskRepository {
	return &UserTask{}
}

/*
 * Default Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaFormKey ...
func (utask *UserTask) SetCamundaFormKey(formKey string) {
	utask.CamundaFormKey = formKey
}

// SetCamundaAsyncBefore ...
func (utask *UserTask) SetCamundaAsyncBefore(asyncBefore bool) {
	utask.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (utask *UserTask) SetCamundaAsyncAfter(asyncAfter bool) {
	utask.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (utask *UserTask) SetCamundaJobPriority(priority int) {
	utask.CamundaJobPriority = priority
}

// SetCamundaAssignee ...
func (utask *UserTask) SetCamundaAssignee(assignee string) {
	utask.CamundaAssignee = assignee
}

// SetCamundaCandidUsers ...
func (ut *UserTask) SetCamundaCandidateUsers(users string) {
	ut.CamundaCandidateUsers = users
}

// SetCamundaCandidGroups ...
func (ut *UserTask) SetCamundaCandidateGroups(groups string) {
	ut.CamundaCandidateGroups = groups
}

// SetCamundaDueDate ...
// rule: due date as an EL expression, e.g. {someDate} or as an ISO date, like 2015-06-29T09:54:00
func (utask *UserTask) SetCamundaDueDate(due string) {
	utask.CamundaDueDate = due
}

// SetCamundaFollowUpDate
// rule: due date as an EL expression, e.g. {someDate} or as an ISO date, like 2015-06-29T09:54:00
func (utask *UserTask) SetCamundaFollowUpDate(followUp string) {
	utask.CamundaFollowUpDate = followUp
}

// SetCamundaPriority ...
func (utask *UserTask) SetCamundaPriority(priority int) {
	utask.CamundaPriority = priority
}

/*** Make Elements ***/

/** BPMN **/

/*** Attributes ***/

// SetExtensionElements ...
func (utask *UserTask) SetExtensionElements() {
	utask.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaFormKey ...
func (utask UserTask) GetCamundaFormKey() impl.STR_PTR {
	return &utask.CamundaFormKey
}

// GetCamundaAsyncBefore ...
func (utask UserTask) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &utask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (utask UserTask) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &utask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (utask UserTask) GetCamundaJobPriority() impl.INT_PTR {
	return &utask.CamundaJobPriority
}

// GetCamundaAssignee ...
func (utask UserTask) GetCamundaAssignee() impl.STR_PTR {
	return &utask.CamundaAssignee
}

// GetCamundaCandidUsers ...
func (ut UserTask) GetCamundaCandidateUsers() impl.STR_PTR {
	return &ut.CamundaCandidateUsers
}

// GetCamundaCandidGroups ...
func (ut UserTask) GetCamundaCandidateGroups() impl.STR_PTR {
	return &ut.CamundaCandidateGroups
}

// GetCamundaDueDate ...
// rule: due date as an EL expression, e.g. {someDate} or as an ISO date, like 2015-06-29T09:54:00
func (utask UserTask) GetCamundaDueDate() impl.STR_PTR {
	return &utask.CamundaDueDate
}

// GetCamundaFollowUpDate
// rule: due date as an EL expression, e.g. {someDate} or as an ISO date, like 2015-06-29T09:54:00
func (utask UserTask) GetCamundaFollowUpDate() impl.STR_PTR {
	return &utask.CamundaFollowUpDate
}

// GetCamundaPriority ...
func (utask UserTask) GetCamundaPriority() impl.INT_PTR {
	return &utask.CamundaPriority
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetExtensionElements ...
func (utask UserTask) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &utask.ExtensionElements[0]
}
