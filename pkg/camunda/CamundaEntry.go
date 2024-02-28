package camunda

import gobpmnTypes "github.com/deemount/gobpmnTypes"

// NewCamundaEntry ...
func NewCamundaEntry() CamundaEntryRepository {
	return &CamundaEntry{}
}

/*
 * @Setters
 */

/* Attributes */

// SetKey ...
func (centry *CamundaEntry) SetKey(key string) {
	centry.Key = key
}

/* Content */

// SetValue ...
func (centry *CamundaEntry) SetValue(value string) {
	centry.Value = value
}

/*
 * @Getters
 */

/* Attributes */

// GetKey ...
func (centry CamundaEntry) GetKey() gobpmnTypes.STR_PTR {
	return &centry.Key
}

/* Content */

// GetValue ...
func (centry CamundaEntry) GetValue() gobpmnTypes.STR_PTR {
	return &centry.Value
}
