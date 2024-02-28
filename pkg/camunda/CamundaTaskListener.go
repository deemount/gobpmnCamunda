package camunda

import gobpmnTypes "github.com/deemount/gobpmnTypes"

// NewCamundaTaskListener ...
func NewCamundaTaskListener() CamundaTaskListenerRepository {
	return &CamundaTaskListener{}
}

/*
 * @Setters
 */

/* Attributes */

// SetClass ...
func (taskListener *CamundaTaskListener) SetClass(class string) {
	taskListener.Class = class
}

// SetEvent ...
// can be: assignment, create, complete, delete, update, timeout
func (taskListener *CamundaTaskListener) SetEvent(event string) {
	taskListener.Event = event
}

// SetID ...
func (taskListener *CamundaTaskListener) SetListenerID(listenerID string) {
	taskListener.ListenerID = listenerID
}

/* Elements */

/** Camunda **/

// SetCamundaField ...
func (taskListener *CamundaTaskListener) SetCamundaField(num int) {
	taskListener.CamundaField = make(CAMUNDA_FIELD_SLC, num)
}

/*
 * @Getters
 */

/* Attributes */

// GetClass ...
func (taskListener CamundaTaskListener) GetClass() gobpmnTypes.STR_PTR {
	return &taskListener.Class
}

// GetEvent ...
// can be: assignment, create, complete, delete, update, timeout
func (taskListener CamundaTaskListener) GetEvent() gobpmnTypes.STR_PTR {
	return &taskListener.Event
}

// GetID ...
func (taskListener CamundaTaskListener) GetListenerID() gobpmnTypes.STR_PTR {
	return &taskListener.ListenerID
}

/* Elements */

/** Camunda **/

// GetCamundaField ...
func (taskListener CamundaTaskListener) GetCamundaField(num int) CAMUNDA_FIELD_PTR {
	return &taskListener.CamundaField[num]
}
