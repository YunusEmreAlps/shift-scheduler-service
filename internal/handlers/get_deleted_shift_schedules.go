package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"shift-scheduler-service/internal/models"
	"shift-scheduler-service/pkg/httpErrors"
)

// HandleGetOnlyDeletedShiftSchedules godoc
// HandleGetOnlyDeletedShiftSchedules handles the request to get all deleted shift schedules
// @Summary get all deleted shift schedules
// @Schemes
// @Description get all deleted shift schedules
// @Tags shift schedules
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} RespondJson "get all shift schedules successfully"
// @Failure 400 {object} RespondJson "cannot get only deleted shift schedules due to invalid request body"
// @Failure 422 {object} RespondJson "cannot get only deleted shift schedules due to invalid request body"
// @Failure 500 {object} RespondJson "cannot get only deleted shift schedules due to internal server error"
// @Router /shift-scheduler-service/shift-schedules/deleted [get]

func (ss *ShiftService) HandleGetOnlyDeletedShiftSchedules(c *gin.Context) (int, interface{}, error) {
	// Step 1: Get all shift schedules from database
	var shift_schedules []models.ShiftSchedule
	if err := ss.db.Unscoped().Where("deleted_at IS NOT NULL").Find(&shift_schedules).Error; err != nil {
		r, i := httpErrors.ErrorResponse(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return r, i, errors.New("cannot get all shift schedules due to not found")
		}
		return r, i, errors.New("cannot get all shift schedules due to internal server error")
	}

	// Step 2: Return all shift schedules
	return http.StatusOK, shift_schedules, nil
}
