package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"shift-scheduler-service/internal/models"
)

// HandleGetShiftScheduleByWeek godoc
// HandleGetShiftScheduleByWeek handles the request to get shift schedules by the current week
// @Summary get shift schedules by current week
// @Schemes
// @Description get shift schedules by current week
// @Tags Shift
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} RespondJson "get shifts by current week successfully"
// @Failure 400 {object} RespondJson "cannot get shifts schedule by current week due to invalid request body"
// @Failure 422 {object} RespondJson "cannot get shifts schedule by current week due to invalid request body"
// @Failure 500 {object} RespondJson "cannot get shifts schedule by current week due to internal server error"
// @Router /shift-schedules/week [get]
func (ss *ShiftService) HandleGetShiftScheduleByWeek(c *gin.Context) (int, interface{}, error) {
	// Step 1: Get all shift schedules
	var shiftSchedules []models.ShiftSchedule
	if err := ss.db.Where("deleted_at IS NULL").Find(&shiftSchedules).Error; err != nil {
		return http.StatusInternalServerError, nil, err
	}

	// Step 2: Get current date and time this week format
	today := time.Now().UTC()
	weekStart := today.AddDate(0, 0, -int(today.Weekday())+1) // Must be Monday
	if today.Weekday() == time.Sunday {
		weekStart = today.AddDate(0, 0, -6) // Adjust for Sunday
	}
	weekStart = time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, time.UTC)
	weekEnd := weekStart.AddDate(0, 0, 7)
	weekEnd = time.Date(weekEnd.Year(), weekEnd.Month(), weekEnd.Day(), 23, 59, 59, 0, time.UTC)

	// Step 3: Filter shift schedules by current week
	var data []map[string]interface{}
	for _, shiftSchedule := range shiftSchedules {
		organizationData := map[string]interface{}{
			"week_start":   weekStart.Format(time.RFC3339),
			"week_end":     weekEnd.Format(time.RFC3339),
			"organization": shiftSchedule.Organization,
			"shifts":       []interface{}{},
		}

		for _, shift := range shiftSchedule.Shifts {
			// Unmarshal shift to map
			shiftMap, ok := shift.(map[string]interface{})
			if !ok {
				continue // Skip if shift is not a map
			}
			startTime, err := time.Parse(time.RFC3339, shiftMap["start"].(string))
			if err != nil {
				continue // Skip invalid date formats
			}
			endTime, err := time.Parse(time.RFC3339, shiftMap["end"].(string))
			if err != nil {
				continue // Skip invalid date formats
			}

			// Check if shift is in current week
			if (startTime.After(weekStart) || startTime.Equal(weekStart)) && (endTime.Before(weekEnd) || endTime.Equal(weekEnd)) {
				organizationData["shifts"] = append(organizationData["shifts"].([]interface{}), shift)
			}
		}

		// Only add organization data if there are shifts in the current week
		if len(organizationData["shifts"].([]interface{})) > 0 {
			data = append(data, organizationData)
		} else {
			data = append(data, map[string]interface{}{
				"week_start":   weekStart.Format(time.RFC3339),
				"week_end":     weekEnd.Format(time.RFC3339),
				"organization": shiftSchedule.Organization,
				"shifts":       []interface{}{},
			})
		}
	}

	// Step 4: Return shift schedules by week
	return http.StatusOK, data, nil
}
