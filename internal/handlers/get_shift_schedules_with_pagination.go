package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"shift-scheduler-service/internal/models"
	"shift-scheduler-service/pkg/httpErrors"
)

// HandleGetShiftSchedulesWithPagination godoc
// HandleGetShiftSchedulesWithPagination handles the request to get all shift schedules with pagination
// @Summary get all shift schedules with pagination
// @Schemes
// @Description get all shift schedules with pagination
// @Tags Shift
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query string false "page"
// @Param size query string false "size"
// @Param orderBy query string false "orderBy"
// @Success 200 {object} RespondJson "get all shift schedules successfully"
// @Failure 400 {object} RespondJson "cannot get shift schedules  with pagination due to invalid request body"
// @Failure 422 {object} RespondJson "cannot get shift schedules  with pagination due to invalid request body"
// @Failure 500 {object} RespondJson "cannot get shift schedules  with pagination due to internal server error"
// @Router /shift-schedules/paginated [get]
func (ss *ShiftService) HandleGetShiftSchedulesWithPagination(c *gin.Context) (int, interface{}, error) {
	// Step 1: Get pagination query params and validate them
	var params models.PaginationQuery
	if err := c.ShouldBindQuery(&params); err != nil {
		return http.StatusBadRequest, nil, err
	}
	// Step 2: Set pagination query params
	params.SetPage(c.Query("page"))
	params.SetSize(c.Query("size"))
	params.SetOrderBy(c.Query("orderBy"))

	// order by validation
	if params.OrderBy != "" {
		if params.OrderBy != "asc" && params.OrderBy != "desc" {
			return http.StatusBadRequest, nil, errors.New("cannot get shifts with pagination due to invalid order by")
		}
		params.OrderBy = fmt.Sprintf("created_at %s", params.OrderBy)
	}

	// Step 3: Get shift with pagination from database
	var shifts []models.ShiftSchedule
	if err := ss.db.Limit(params.Size).Offset(params.GetOffset()).Order(params.OrderBy).Find(&shifts).Error; err != nil {
		r, i := httpErrors.ErrorResponse(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return r, i, errors.New("cannot get shifts with pagination due to not found")
		}
		return r, i, errors.New("cannot get shifts with pagination due to internal server error")
	}

	// Step 4: Return shifts with pagination
	return http.StatusOK, shifts, nil
}
