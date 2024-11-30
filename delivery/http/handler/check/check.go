package checkhandler

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/ruhollahh/go-ecom/internal/clients/dbpostgre"
	"net/http"
	"os"
	"runtime"
	"time"
)

// Readiness godoc
// @Summary 	 Check readiness
// @Tags         Checks
// @Accept       json
// @Produce      json
// @Success      200  {object} DlvReadinessResp
// @Failure      400  {object} expectederr.ErrorDocument{error=string,fields=nil}
// @Failure      500  {object} expectederr.ErrorDocument{error=string,fields=nil}
// @Router       /v1/checks/readiness [get].
func (h *Handler) Readiness(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second)
	defer cancel()

	status := "ok"
	statusCode := http.StatusOK
	if err := dbpostgre.StatusCheck(ctx, h.db); err != nil {
		status = "db not ready"
		statusCode = http.StatusInternalServerError
	}

	data := DlvReadinessResp{
		Status: status,
	}

	return c.JSON(statusCode, data)
}

// Liveness godoc
// @Summary 	 Check liveness
// @Tags         Checks
// @Accept       json
// @Produce      json
// @Success      200  {object} DlvLivenessResp
// @Failure      400  {object} expectederr.ErrorDocument{error=string}
// @Failure      500  {object} expectederr.ErrorDocument{error=string}
// @Router       /v1/checks/liveness [get].
func (h *Handler) Liveness(c echo.Context) error {
	host, err := os.Hostname()
	if err != nil {
		host = "unavailable"
	}

	data := DlvLivenessResp{
		Status:     "up",
		Build:      h.build,
		Host:       host,
		GOMAXPROCS: runtime.GOMAXPROCS(0),
	}

	return c.JSON(http.StatusOK, data)
}
