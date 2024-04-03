package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"shift-scheduler-service/internal/models"
	"shift-scheduler-service/pkg/httpErrors"
)

// HandleRestoreShiftSchedule godoc
// HandleRestoreShiftSchedule handles the request to restore a shift schedule
// @Summary restore a shift schedule
// @Schemes
// @Description restored a shift schedule
// @Tags shift schedule
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} RespondJson "successfully restored shift schedule"
// @Failure 400 {object} RespondJson "cannot restore shift schedule due to invalid request body"
// @Failure 422 {object} RespondJson "cannot restore shift schedule due to invalid request body"
// @Failure 500 {object} RespondJson "cannot restore shift schedule due to internal server error"
// @Router /shift-scheduler-service/shift-schedules/{id}/restore [restore]

func (ss *ShiftService) HandleRestoreShiftSchedule(c *gin.Context) (int, interface{}, error) {
	// Step 1: Get shift schedule id from path	and validate
	id := c.Param("id")
	if id == "" {
		return http.StatusBadRequest, nil, nil
	}

	// Step 2: Restore shift schedule by id from database
	var shiftSchedule models.ShiftSchedule
	if err := ss.db.Unscoped().Model(shiftSchedule).Where("id = ?", id).Update("deleted_at", nil).Error; err != nil {
		r, i := httpErrors.ErrorResponse(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return r, i, errors.New("cannot restore shift schedule due to not found")
		}
		return r, i, errors.New("cannot restore shift schedule due to internal server error")
	}

	// Step 3: Return shift schedule by id
	return http.StatusOK, "Shift Schedule Successfully Restored", nil
}
