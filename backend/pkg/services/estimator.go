package services

// EstimateSalary returns a dummy salary estimate based on the role, location, experience, and external data.
func EstimateSalary(role string, location string, experience int, externalData []int) int {
	baseSalary := 60000
	experienceBonus := experience * 2000

	externalAdjustment := 0
	if len(externalData) > 0 {
		sum := 0
		for _, v := range externalData {
			sum += v
		}
		externalAdjustment = sum / len(externalData)
	}

	return baseSalary + experienceBonus + externalAdjustment
}
