package core_http_response

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_logger "github.com/haytoku/golang-todoapp/internal/core/logger"
	"go.uber.org/zap"
)

type HHTPResponseHandler struct {
	log *core_logger.Logger
	rw  http.ResponseWriter
}

func NewHTTPResponseHandler(log *core_logger.Logger, rw http.ResponseWriter) *HHTPResponseHandler {
	return &HHTPResponseHandler{
		log: log,
		rw:  rw,
	}
}

func (h *HHTPResponseHandler) PanicResponse(p any, msg string) {

	statusCode := http.StatusInternalServerError

	err := fmt.Errorf("unexpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))
	h.rw.WriteHeader(statusCode)

	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	if err := json.NewEncoder(h.rw).Encode(response); err != nil {
		h.log.Error("failed to encode response", zap.Error(err))
	}

}
