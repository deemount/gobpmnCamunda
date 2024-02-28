package tasks

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewUserTask ...
func NewUserTask() UserTaskRepository {
	return &UserTask{}
}

/*
 * @Setters
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
	utask.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaFormKey ...
func (utask UserTask) GetCamundaFormKey() gobpmnTypes.STR_PTR {
	return &utask.CamundaFormKey
}

// GetCamundaAsyncBefore ...
func (utask UserTask) GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR {
	return &utask.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (utask UserTask) GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR {
	return &utask.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (utask UserTask) GetCamundaJobPriority() gobpmnTypes.INT_PTR {
	return &utask.CamundaJobPriority
}

// GetCamundaAssignee ...
func (utask UserTask) GetCamundaAssignee() gobpmnTypes.STR_PTR {
	return &utask.CamundaAssignee
}

// GetCamundaCandidUsers ...
func (ut UserTask) GetCamundaCandidateUsers() gobpmnTypes.STR_PTR {
	return &ut.CamundaCandidateUsers
}

// GetCamundaCandidGroups ...
func (ut UserTask) GetCamundaCandidateGroups() gobpmnTypes.STR_PTR {
	return &ut.CamundaCandidateGroups
}

// GetCamundaDueDate ...
// rule: due date as an EL expression, e.g. {someDate} or as an ISO date, like 2015-06-29T09:54:00
func (utask UserTask) GetCamundaDueDate() gobpmnTypes.STR_PTR {
	return &utask.CamundaDueDate
}

// GetCamundaFollowUpDate
// rule: due date as an EL expression, e.g. {someDate} or as an ISO date, like 2015-06-29T09:54:00
func (utask UserTask) GetCamundaFollowUpDate() gobpmnTypes.STR_PTR {
	return &utask.CamundaFollowUpDate
}

// GetCamundaPriority ...
func (utask UserTask) GetCamundaPriority() gobpmnTypes.INT_PTR {
	return &utask.CamundaPriority
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// GetExtensionElements ...
func (utask UserTask) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &utask.ExtensionElements[0]
}
