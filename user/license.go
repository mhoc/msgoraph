package user

// AssignedLicense represents a license assigned to a user
type AssignedLicense struct {
	DisabledPlans []string `json:"disabledPlans"`
	SKUID         *string  `json:"skuId"`
}
