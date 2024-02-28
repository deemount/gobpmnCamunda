package tasks

import (
	"github.com/deemount/gobpmnCamunda/pkg/camunda"
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
)

type Tasks struct {
	BusinessRuleTask BUSINESS_RULE_TASK_SLC `xml:"bpmn:businessRuleTask,omitempty" json:"businessRuleTask,omitempty" csv:"-"`
	Task             TASK_SLC               `xml:"bpmn:task,omitempty" json:"task,omitempty" csv:"-"`
	UserTask         USER_TASK_SLC          `xml:"bpmn:userTask,omitempty" json:"userTask,omitempty" csv:"-"`
	ManualTask       MANUAL_TASK_SLC        `xml:"bpmn:manualTask,omitempty" json:"manualTask,omitempty" csv:"-"`
	ReceiveTask      RECEIVE_TASK_SLC       `xml:"bpmn:receiveTask,omitempty" json:"receiveTask,omitempty" csv:"-"`
	ScriptTask       SCRIPT_TASK_SLC        `xml:"bpmn:scriptTask,omitempty" json:"scriptTask,omitempty" csv:"-"`
	SendTask         SEND_TASK_SLC          `xml:"bpmn:sendTask,omitempty" json:"sendTask,omitempty" csv:"-"`
	ServiceTask      SERVICE_TASK_SLC       `xml:"bpmn:serviceTask,omitempty" json:"serviceTask,omitempty" csv:"-"`
}

type TTasks struct {
	BusinessRuleTask BUSINESS_RULE_TASK_SLC `xml:"businessRuleTask,omitempty" json:"businessRuleTask,omitempty" csv:"-"`
	Task             TASK_SLC               `xml:"task,omitempty" json:"task,omitempty" csv:"-"`
	UserTask         USER_TASK_SLC          `xml:"userTask,omitempty" json:"userTask,omitempty" csv:"-"`
	ManualTask       MANUAL_TASK_SLC        `xml:"manualTask,omitempty" json:"manualTask,omitempty" csv:"-"`
	ReceiveTask      RECEIVE_TASK_SLC       `xml:"receiveTask,omitempty" json:"receiveTask,omitempty" csv:"-"`
	ScriptTask       SCRIPT_TASK_SLC        `xml:"scriptTask,omitempty" json:"scriptTask,omitempty" csv:"-"`
	SendTask         SEND_TASK_SLC          `xml:"sendTask,omitempty" json:"sendTask,omitempty" csv:"-"`
	ServiceTask      SERVICE_TASK_SLC       `xml:"serviceTask,omitempty" json:"serviceTask,omitempty" csv:"-"`
}

/*
 * @Elementary
 */

// BusinessRuleTask ...
type BusinessRuleTask struct {
	extension_elements.Extension_Elements
	camunda.CoreAttributes
	CamundaClass string `xml:"camunda:class,attr,omitempty" json:"class,omitempty"`
}

// TBusinessRuleTask ...
type TBusinessRuleTask struct {
	extension_elements.TExtension_Elements
	camunda.TCoreAttributes
	Class string `xml:"class,attr,omitempty" json:"class,omitempty"`
}

// ManualTask ...
type ManualTask struct {
	extension_elements.Extension_Elements
	camunda.CoreAttributes
}

// TManualTask ...
type TManualTask struct {
	extension_elements.TExtension_Elements
	camunda.TCoreAttributes
}

// ReceiveTask ...
type ReceiveTask struct {
	extension_elements.Extension_Elements
	camunda.CoreAttributes
}

// TReceiveTask ...
type TReceiveTask struct {
	extension_elements.TExtension_Elements
	camunda.TCoreAttributes
}

// ScriptTask ...
type ScriptTask struct {
	extension_elements.Extension_Elements
	camunda.CoreAttributes
}

// TScriptTask ...
type TScriptTask struct {
	extension_elements.TExtension_Elements
	camunda.TCoreAttributes
}

// SendTask ...
type SendTask struct {
	extension_elements.Extension_Elements
	camunda.CoreAttributes
}

// TSendTask ...
type TSendTask struct {
	extension_elements.TExtension_Elements
	camunda.TCoreAttributes
}

// ServiceTask ...
type ServiceTask struct {
	extension_elements.Extension_Elements
	camunda.CoreAttributes
}

// TServiceTask ...
type TServiceTask struct {
	extension_elements.TExtension_Elements
	camunda.TCoreAttributes
}

// Task ...
type Task struct {
	extension_elements.Extension_Elements
	camunda.CoreAttributes
}

// TTask ...
type TTask struct {
	extension_elements.TExtension_Elements
	camunda.TCoreAttributes
}

// UserTask ...
type UserTask struct {
	extension_elements.Extension_Elements
	camunda.CoreAttributes
	CamundaFormKey         string `xml:"camunda:formKey,attr,omitempty" json:"formKey,omitempty"`
	CamundaAssignee        string `xml:"camunda:assignee,attr,omitempty" json:"assignee,omitempty"`
	CamundaCandidateUsers  string `xml:"camunda:candidateUsers,attr,omitempty" json:"candidateUsers,omitempty"`
	CamundaCandidateGroups string `xml:"camunda:candidateGroups,attr,omitempty" json:"candidateGroups,omitempty"`
	CamundaDueDate         string `xml:"camunda:dueDate,attr,omitempty" json:"dueDate,omitempty"`
	CamundaFollowUpDate    string `xml:"camunda:followUpDate,attr,omitempty" json:"followUpDate,omitempty"`
	CamundaPriority        int    `xml:"camunda:priority,attr,omitempty" json:"priority,omitempty"`
}

// TUserTask ...
type TUserTask struct {
	extension_elements.TExtension_Elements
	camunda.TCoreAttributes
	FormKey         string `xml:"formKey,attr,omitempty" json:"formKey,omitempty"`
	Assignee        string `xml:"assignee,attr,omitempty" json:"assignee,omitempty"`
	CandidateUsers  string `xml:"candidateUsers,attr,omitempty" json:"candidateUsers,omitempty"`
	CandidateGroups string `xml:"candidateGroups,attr,omitempty" json:"candidateGroups,omitempty"`
	DueDate         string `xml:"dueDate,attr,omitempty" json:"dueDate,omitempty"`
	FollowUpDate    string `xml:"followUpDate,attr,omitempty" json:"followUpDate,omitempty"`
	Priority        int    `xml:"priority,attr,omitempty" json:"priority,omitempty"`
}
