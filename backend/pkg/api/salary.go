// pkg/api/salary.go
package api

import (
	"encoding/json"
	"net/http"

	"github.com/camscott16/wip/pkg/services"
)

// SalaryRequest represents the expected request body.
type SalaryRequest struct {
	Role       string `json:"role"`
	Location   string `json:"location"`
	Experience int    `json:"experience"`
}

// SalaryResponse represents the JSON response.
type SalaryResponse struct {
	Role            string `json:"role"`
	Location        string `json:"location"`
	Experience      int    `json:"experience"`
	EstimatedSalary int    `json:"estimatedSalary"`
}

// SalaryEstimateHandler handles POST requests for salary estimation.
func SalaryEstimateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SalaryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// For now, externalData is an empty slice; later, you can integrate scraped data.
	externalData := []int{}
	estimatedSalary := services.EstimateSalary(req.Role, req.Location, req.Experience, externalData)

	res := SalaryResponse{
		Role:            req.Role,
		Location:        req.Location,
		Experience:      req.Experience,
		EstimatedSalary: estimatedSalary,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
