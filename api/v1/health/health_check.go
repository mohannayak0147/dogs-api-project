package health

import (
	"dogs-api/version"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

var upSince = time.Now()

type healthCheck struct {
	Alive     bool   `json:"alive"`
	Since     string `json:"since"`
	Version   string `json:"version"`
	GoVersion string `json:"go_version"`
}

// HealthCheckHandler : A very simple health check.
// @Summary      Health Check
// @Description  Returns the status of the running application
// @Tags         HealthCheck
// @Produce      json
// @Success      200  {object}  healthCheck  "Global health of the application"
// @Failure      500  {object}  object       "Error description formated as {"msg":"string"}"
// @Router       /api/status [get]
func HealthCheckHandler(c *fiber.Ctx) error {

	log.Debug("handle health check")
	healthCheck := &healthCheck{
		Alive:     true,
		Since:     time.Since(upSince).String(),
		Version:   version.Version,
		GoVersion: version.GoVersion,
	}
	responseBytes, err := json.Marshal(healthCheck)
	if err != nil {
		log.Error("unable to marshall healthCheck response: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})
	}
	return c.Send(responseBytes)
}
