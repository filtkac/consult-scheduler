package integrationtest

import (
	"bytes"
	"consult-scheduler/config"
	"consult-scheduler/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDepartmentControllerFlow(t *testing.T) {
	tx := testDB.Begin()
	defer tx.Rollback()
	router := config.RouterConfig(tx)

	// Get departments
	departments := getDepartments(t, router)

	assert.Len(t, departments, 2)
	assert.Equal(t, "Mammo center", departments[0].Name)
	assert.Equal(t, "CT surgery", departments[1].Name)

	// Delete non-existing department, nothing happens
	req := httptest.NewRequest(http.MethodDelete, "/departments/999", nil)
	w := makeRequest(router, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

	assert.Len(t, getDepartments(t, router), 2)

	// Delete an existing department, check that the correct one was removed
	req = httptest.NewRequest(http.MethodDelete, "/departments/1", nil)
	makeRequest(router, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

	departments = getDepartments(t, router)
	assert.Len(t, departments, 1)
	assert.Equal(t, "CT surgery", departments[0].Name)

	// Create a new department
	departmentToCreate := model.DepartmentDto{
		Name: "Cardiology",
	}
	body, err := json.Marshal(departmentToCreate)
	assert.NoError(t, err)

	req = httptest.NewRequest(http.MethodPost, "/departments", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = makeRequest(router, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	var newDepartment model.DepartmentDto
	err = json.Unmarshal(w.Body.Bytes(), &newDepartment)
	assert.NoError(t, err)

	assert.Equal(t, "Cardiology", newDepartment.Name)

	// Get all departments again and check that the new one was added
	departments = getDepartments(t, router)
	assert.Len(t, departments, 2)
	assert.Equal(t, "CT surgery", departments[0].Name)
	assert.Equal(t, "Cardiology", departments[1].Name)
}

func getDepartments(t *testing.T, router *gin.Engine) []model.DepartmentDto {
	req := httptest.NewRequest(http.MethodGet, "/departments", nil)
	w := makeRequest(router, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var existingDepartments []model.DepartmentDto
	err := json.Unmarshal(w.Body.Bytes(), &existingDepartments)
	assert.NoError(t, err)
	return existingDepartments
}
