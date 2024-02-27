package subprocesses

import (
	"github.com/deemount/gobpmnCamunda/pkg/attributes"
	impl "github.com/deemount/gobpmnTypes"
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
	transaction.ExtensionElements = make([]attributes.ExtensionElements, 1)
}

/*
 * @Getters
 */

/* Attributes */

/** Camunda **/

// GetCamundaAsyncBefore ...
func (transaction Transaction) GetCamundaAsyncBefore() impl.BOOL_PTR {
	return &transaction.CamundaAsyncBefore
}

// GetCamundaAsyncAfter ...
func (transaction Transaction) GetCamundaAsyncAfter() impl.BOOL_PTR {
	return &transaction.CamundaAsyncAfter
}

// GetCamundaJobPriority ...
func (transaction Transaction) GetCamundaJobPriority() impl.INT_PTR {
	return &transaction.CamundaJobPriority
}

/* Elements */

/** BPMN **/

// GetExtensionElements ...
func (transaction Transaction) GetExtensionElements() attributes.EXTENSION_ELEMENTS_PTR {
	return &transaction.ExtensionElements[0]
}
