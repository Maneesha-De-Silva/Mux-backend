package employee

import (
	"backend/internal/types"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// get all employees
func (h *Handler) FetchEmployee(w http.ResponseWriter, r *http.Request) {
	employees := h.EmployeeService.GetPosts()
	h.HandleSuccessRespose(w, employees)
}

// create a new employee
func (h *Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var newEmployee types.Employee

	if err := json.NewDecoder(r.Body).Decode(&newEmployee); err != nil {
		h.HandleErrorRespose(w, "Failed to decode JSON body", err, http.StatusInternalServerError)
		return
	}

	employee, err := h.EmployeeService.SaveEmployee(newEmployee)

	if err != nil {
		h.HandleErrorRespose(w, "Unable to save the post", err, http.StatusBadRequest)
		return
	}

	h.HandleSuccessRespose(w, employee)
}

// get item details by ID
func (h *Handler) GetEmployeeDetalis(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		h.HandleErrorRespose(w, "Unable to pass int", err, http.StatusBadRequest)
		return
	}

	details := h.EmployeeService.GetEmployeeByID(i)

	h.HandleSuccessRespose(w, details)

}

// delete a item
func (h *Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		h.HandleErrorRespose(w, "Unable to pass int", err, http.StatusBadRequest)
		return
	}

	h.EmployeeService.DeleteEmployee(i)
	h.HandleSuccessRespose(w, struct {
		Message string `json:"message"`
	}{
		Message: "Success",
	})
}
