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

type createShiftScheduleDTO struct {
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

// HandleCreateShiftSchedule godoc
// HandleCreateShiftSchedule handles the request to create a new shift schedule
// @Summary create a new shift schedule
// @Schemes
// @Description create a new shift schedule
// @Tags shift schedule
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} RespondJson "successfully created shift schedule"
// @Failure 400 {object} RespondJson "cannot create shift schedule due to invalid request body"
// @Failure 422 {object} RespondJson "cannot create shift schedule due to invalid request body"
// @Failure 500 {object} RespondJson "cannot create shift schedule due to internal server error"
// @Router /shift-scheduler-service/shift-schedules [post]

func (ss *ShiftService) HandleCreateShiftSchedule(c *gin.Context) (int, interface{}, error) {
	// Step 1: Get shift schedule from request body
	var params createShiftScheduleDTO
	if err := c.ShouldBindJSON(&params); err != nil {
		return http.StatusBadRequest, nil, err
	}
	// if err := utils.ValidateStruct(c, &params); err != nil {
	// 	return http.StatusBadRequest, nil, err
	// }

	// Step 2: Validate shift schedule
	var shift_schedule models.ShiftSchedule
	createParamsToShiftSchedule(&params, &shift_schedule)

	// Step 3: Create shift schedule in database
	if err := ss.db.Create(&shift_schedule).Error; err != nil {
		r, i := httpErrors.ErrorResponse(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return r, i, errors.New("cannot create shift schedule due to not found")
		}
		return r, i, errors.New("cannot create shift schedule due to internal server error")
	}

	// Step 4: Return shift schedule
	return http.StatusOK, "Shift Schedule Successfully Created", nil
}

func createParamsToShiftSchedule(params *createShiftScheduleDTO, shift *models.ShiftSchedule) {
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
