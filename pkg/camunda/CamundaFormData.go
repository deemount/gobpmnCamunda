package camunda

// NewCamundaFormData ...
func NewCamundaFormData() CamundaFormDataRepository {
	return &CamundaFormData{}
}

/*
 * @Setters
 */

/* Elements */

/** Camunda **/

// SetCamundaFormField ...
func (formData *CamundaFormData) SetCamundaFormField(num int) {
	formData.CamundaFormField = make(CAMUNDA_FORM_FIELD_SLC, num)
}

/*
 * @Getters
 */

/* Elements */

/** Camunda **/

// GetCamundaFormField ...
func (formData CamundaFormData) GetCamundaFormField(num int) CAMUNDA_FORM_FIELD_PTR {
	return &formData.CamundaFormField[num]
}
