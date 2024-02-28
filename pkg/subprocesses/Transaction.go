package subprocesses

import (
	"github.com/deemount/gobpmnCamunda/pkg/extension_elements"
	gobpmnTypes "github.com/deemount/gobpmnTypes"
)

// NewTransaction ...
func NewTransaction() TransactionRepository {
	return &Transaction{}
}

/*
 * @Setters
 */

/* Attributes */

/** Camunda **/

// SetCamundaAsyncBefore ...
func (transaction *Transaction) SetCamundaAsyncBefore(asyncBefore bool) {
	transaction.CamundaAsyncBefore = asyncBefore
}

// SetCamundaAsyncAfter ...
func (transaction *Transaction) SetCamundaAsyncAfter(asyncAfter bool) {
	transaction.CamundaAsyncAfter = asyncAfter
}

// SetCamundaJobPriority ...
func (transaction *Transaction) SetCamundaJobPriority(priority int) {
	transaction.CamundaJobPriority = priority
}

/* Elements */

/** BPMN **/

/*** Attributes ***/

// SetExtensionElements ...
func (transaction *Transaction) SetExtensionElements() {
	transaction.ExtensionElements = make(extension_elements.EXTENSION_ELEMENTS_SLC, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (transaction Transaction) GetCamundaAsyncBefore() gobpmnTypes.BOOL_PTR {
	return &transaction.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (transaction Transaction) GetCamundaAsyncAfter() gobpmnTypes.BOOL_PTR {
	return &transaction.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (transaction Transaction) GetCamundaJobPriority() gobpmnTypes.INT_PTR {
	return &transaction.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (transaction Transaction) GetExtensionElements() extension_elements.EXTENSION_ELEMENTS_PTR {
	return &transaction.ExtensionElements[0]
}
