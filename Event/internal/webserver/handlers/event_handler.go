package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"Projeto-booster/internal/dto"
	"Projeto-booster/internal/entity"
	"Projeto-booster/internal/infra/database"
)

type EventHandler struct {
	EventDB database.EventInterface
}

func NewEventHandler(eventDB database.EventInterface) *EventHandler {
	return &EventHandler{EventDB: eventDB}
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event dto.CreateEventInput
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	e, err := entity.NewEvent(event.Name, event.Description, event.StartDate, event.FinishDate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.EventDB.Create(e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

func (h *EventHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	event, err := h.EventDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}

func (h *EventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := h.EventDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var event entity.Event
	err = json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.EventDB.Update(&event)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&event)
}

func (h *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := h.EventDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.EventDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}