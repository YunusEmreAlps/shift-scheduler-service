package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"shift-scheduler-service/internal/models"
	"shift-scheduler-service/pkg/httpErrors"
)

// HandleGetShiftScheduleByYear godoc
// HandleGetShiftScheduleByYear handles the request to get a shift schedule by year
// @Summary get a shift schedule by year
// @Schemes
// @Description get a shift schedule by year
// @Tags shift schedule
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} RespondJson "get shift by year successfully"
// @Failure 400 {object} RespondJson "cannot get shift schedule by year due to invalid request body"
// @Failure 422 {object} RespondJson "cannot get shift schedule by year due to invalid request body"
// @Failure 500 {object} RespondJson "cannot get shift schedule by year due to internal server error"
// @Router /shift-scheduler-service/shift-schedules/id [get]

func (ss *ShiftService) HandleGetShiftScheduleByYear(c *gin.Context) (int, interface{}, error) {
	// Step 1: Get shift year from path and validate
	year := c.Param("year")
	if year == "" {
		return http.StatusBadRequest, nil, nil
	}

	// Step 2: Get shift schedules by year from database
	var shift_schedules models.ShiftSchedule
	if err := ss.db.Where("deleted_at IS NULL AND year = ?", year).Find(&shift_schedules).Error; err != nil {
		r, i := httpErrors.ErrorResponse(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return r, i, errors.New("cannot get shift schedules by year due to not found")
		}
		return r, i, errors.New("cannot get shift schedule by year due to internal server error")
	}

	// Step 3: Return shift schedules by year
	return http.StatusOK, shift_schedules, nil
}
