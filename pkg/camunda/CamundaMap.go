package camunda

// NewCamundaMap ...
func NewCamundaMap() CamundaMapRepository {
	return &CamundaMap{}
}

/*
 * @Setters
 */

/* Elements */

/** Camunda **/

// SetCamundaEntry
func (mp *CamundaMap) SetCamundaEntry(num int) {
	mp.CamundaEntry = make(CAMUNDA_ENTRY_SLC, num)
}

/*
 * @Getters
 */

/* Elements */

/** Camunda **/

// GetCamundaEntry
func (mp CamundaMap) GetCamundaEntry(num int) CAMUNDA_ENTRY_PTR {
	return &mp.CamundaEntry[num]
}
