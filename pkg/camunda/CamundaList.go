package camunda

// NewCamundaList ...
func NewCamundaList() CamundaListRepository {
	return &CamundaList{}
}

/*
 * @Setters
 */

/* Elements */

/** Camunda **/

// SetCamundaValue ...
func (list *CamundaList) SetCamundaValue(num int) {
	list.CamundaValue = make(CAMUNDA_VALUE_SLC, num)
}

/*
 * @Getters
 */

/* Elements */

/** Camunda **/

// GetCamundaValue ...
func (list CamundaList) GetCamundaValue(num int) CAMUNDA_VALUE_PTR {
	return &list.CamundaValue[num]
}
