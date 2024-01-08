package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"shift-scheduler-service/internal/models"
	"shift-scheduler-service/pkg/httpErrors"
)

type updateShiftScheduleDTO struct {
	Alias        string       `json:"alias" binding:"required"`
	Organization models.JSONB `json:"organization" binding:"required"`
	Manager      models.JSONB `json:"manager" binding:"required"`
	Description  string       `json:"description" binding:"required"`
	Start_Date   time.Time    `json:"start_date" binding:"required"`
	End_Date     time.Time    `json:"end_date" binding:"required"`
	Shifts       models.JSONB `json:"shifts" binding:"required"`
	Year         int          `json:"year" binding:"required"`
	Status       int          `json:"status"` // 0: pending, 1: approved, 2: rejected
}

// HandleUpdateShiftSchedule godoc
// HandleUpdateShiftSchedule handles the request to update a shift schedule
// @Summary update a shift schedule
// @Schemes
// @Description update a shift schedule
// @Tags shift schedule
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} RespondJson "successfully updated shift schedule"
// @Failure 400 {object} RespondJson "cannot update shift schedule due to invalid request body"
// @Failure 422 {object} RespondJson "cannot update shift schedule due to invalid request body"
// @Failure 500 {object} RespondJson "cannot update shift schedule due to internal server error"
// @Router /shift-scheduler-service/budget-proposals/{id} [put]

func (ss *ShiftService) HandleUpdateShiftSchedule(c *gin.Context) (int, interface{}, error) {
	// Step 1: Get shift id from path and validate it
	id := c.Param("id")
	if id == "" {
		return http.StatusBadRequest, nil, nil
	}

	// Step 2: Get DTO from request body and validate it
	var params updateShiftScheduleDTO
	if err := c.ShouldBindJSON(&params); err != nil {
		return http.StatusBadRequest, nil, err
	}

	var shift models.ShiftSchedule
	if err := ss.db.Where("id = ?", id).First(&shift).Error; err != nil {
		r, i := httpErrors.ErrorResponse(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return r, i, errors.New("cannot update shift due to not found")
		}
		return r, i, errors.New("cannot update shift due to internal server error")
	}

	// Step 3: Map DTO to shift and validate it
	updateParamsToShiftSchedule(&params, &shift)

	// Step 4: Update shift to database
	if err := ss.db.Save(&shift).Error; err != nil {
		r, i := httpErrors.ErrorResponse(err)
		return r, i, errors.New("cannot update shift due to internal server error")
	}

	// Step 5: Get shift by id from database
	return http.StatusOK, "Shift Schedule Successfully Updated", nil
}

func updateParamsToShiftSchedule(params *updateShiftScheduleDTO, shift *models.ShiftSchedule) {
	shift.Alias = params.Alias
	shift.Organization = params.Organization
	shift.Manager = params.Manager
	shift.Description = params.Description
	shift.Start_Date = params.Start_Date
	shift.End_Date = params.End_Date
	shift.Shifts = params.Shifts
	shift.Year = params.Year
	shift.Status = params.Status
}
