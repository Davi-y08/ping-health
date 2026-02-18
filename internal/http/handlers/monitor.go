package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	appMonitor "ping-health/internal/application/monitor"
	mService "ping-health/internal/application/monitor"
	mapErrors "ping-health/internal/http/http_errors"
	"ping-health/internal/httpx"

	"github.com/google/uuid"
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

	userIDValue := r.Context().Value("user_id")

	if userIDValue == nil {
		return httpx.Unauthorized(errors.New("user not found in context"))
	}

	userIDString, ok := userIDValue.(string)
	if !ok {
		return httpx.Internal(errors.New("invalid user id type"))
	}

	userUUID, err := uuid.Parse(userIDString)
	if err != nil {
		return httpx.Internal(errors.New("invalid user id format"))
	}

	dto.UserID = userUUID

	if create_err := h.service.CreateMonitorService(r.Context(), dto); create_err != nil{
		return mapErrors.MapErrorsMonitor(create_err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("monitor created"))
	return nil
}