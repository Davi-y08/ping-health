package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	appMonitor "ping-health/internal/application/monitor"
	mService "ping-health/internal/application/monitor"
	mapErrors "ping-health/internal/http/http_errors"
	"ping-health/internal/httpx"
)

type MonitorHandler struct{
	service *mService.MonitorService
}

func NewMonitorHandler(service *mService.MonitorService) *MonitorHandler {
	return &MonitorHandler{service: service}
}

func (h *MonitorHandler) CreateMonitorHandler(w http.ResponseWriter, r *http.Request) (error){
	if r.Method != http.MethodPost {
		return httpx.MethodNotAllowed(errors.New("method not allowed"))
	}

	var dto appMonitor.CreateMonitorDto

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return httpx.BadRequest(errors.New("corpo inválido"))
	}

	if create_err := h.service.CreateMonitorService(r.Context(), dto); create_err != nil{
		return mapErrors.MapErrorsMonitor(create_err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("monitor created"))
	return nil
}