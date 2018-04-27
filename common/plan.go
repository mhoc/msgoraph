package common

// AssignedPlan property of both user entity and organization entity is a
// collection of assignedPlan
type AssignedPlan struct {
	AssignedDateTime string `json:"assignedDateTime"`
	CapabilityStatus string `json:"capabilityStatus"`
	Service          string `json:"service"`
	ServicePlanID    string `json:"servicePlanId"`
}

// ProvisionedPlan The provisionedPlans property of the user entity and the
// organization entity is a collection of provisionedPlan.
type ProvisionedPlan struct {
	CapabilityStatus   string `json:"capabilityStatus"`
	ProvisioningStatus string `json:"provisioningStatus"`
	Service            string `json:"service"`
}
