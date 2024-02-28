package camunda

// NewCamundaProperties ...
func NewCamundaProperties() CamundaPropertiesRepository {
	return &CamundaProperties{}
}

/*
 * @Setters
 */

/* Elements */

/** Camunda **/

// SetCamundaProperties ...
func (properties *CamundaProperties) SetCamundaProperty(num int) {
	properties.CamundaProperty = make(CAMUNDA_PROPERTY_SLC, num)
}

/*
 * @Getters
 */

/* Elements */

/** Camunda **/

// GetCamundaProperties ...
func (properties CamundaProperties) GetCamundaProperty(num int) CAMUNDA_PROPERTY_PTR {
	return &properties.CamundaProperty[num]
}
