package handlers

import (
	"breakfast/internal/models"
	"breakfast/internal/services"
	u "breakfast/internal/utilities"
	"strconv"

	"encoding/json"
	"net/http"
)

type OmeletteHandler struct {
	Omelette *services.OmeletteService
}

func NewOmeletteHandler(service *services.OmeletteService) *OmeletteHandler {
	return &OmeletteHandler{Omelette: service}
}

func (h *OmeletteHandler) Create(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	var data models.OmeletteTable
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, err.Error(), nil, http.StatusConflict)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	err = h.Omelette.Create(user_id, &data)
	if err != nil {
		u.Send(w, "Could not create event: "+err.Error(), data, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", data, http.StatusCreated)
}

func (h *OmeletteHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		u.Send(w, "Error reading the ID requested", nil, http.StatusBadRequest)
		return
	}

	data, err := h.Omelette.GetByID(user_id, id)
	if err != nil {
		u.Send(w, "Error retrieving table:"+err.Error(), data, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", data, http.StatusOK)
}

func (h *OmeletteHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	data, err := h.Omelette.GetAll(user_id)
	if err != nil {
		u.Send(w, "Error retrieving event:"+err.Error(), data, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", data, http.StatusOK)
}

func (h *OmeletteHandler) Update(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	var data models.OmeletteTable
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, err.Error(), nil, http.StatusConflict)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		u.Send(w, "Error reading the ID requested", nil, http.StatusBadRequest)
		return
	}

	data.TableID = id
	err, table := h.Omelette.Update(user_id, &data)
	if err != nil {
		u.Send(w, "Error updating table:"+err.Error(), table, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", table, http.StatusOK)
}

func (h *OmeletteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		u.Send(w, "Error reading the ID requested", nil, http.StatusBadRequest)
		return
	}

	err = h.Omelette.Delete(user_id, id)
	if err != nil {
		u.Send(w, "Error retrieving table:"+err.Error(), nil, http.StatusInternalServerError)
		return
	}

	u.Send(w, "DELETED", nil, http.StatusOK)
}

func (h *OmeletteHandler) CreateList(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		u.Send(w, "Error reading the ID requested", nil, http.StatusBadRequest)
		return
	}

	var data models.OmeletteList
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, err.Error(), nil, http.StatusConflict)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	data.TableID = id
	list, err := h.Omelette.CreateList(user_id, id, &data)
	if err != nil {
		u.Send(w, "Could not create list: "+err.Error(), list, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", list, http.StatusCreated)
}

func (h *OmeletteHandler) GetListByID(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	listIDStr := r.PathValue("list_id")
	listID, err := strconv.Atoi(listIDStr)
	if err != nil {
		u.Send(w, "Invalid reminder ID", nil, http.StatusBadRequest)
		return
	}

	list, err := h.Omelette.GetListByID(userID, listID)
	if err != nil {
		u.Send(w, "Could not retrieve list: "+err.Error(), nil, http.StatusNotFound)
		return
	}

	u.Send(w, "", list, http.StatusOK)
}

func (h *OmeletteHandler) GetAllLists(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	lists, err := h.Omelette.GetAllLists(userID)
	if err != nil {
		u.Send(w, "Could not retrieve lists: "+err.Error(), nil, http.StatusNotFound)
		return
	}

	u.Send(w, "", lists, http.StatusOK)
}

func (h *OmeletteHandler) UpdateList(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	tableIDStr := r.PathValue("id")
	tableID, err := strconv.Atoi(tableIDStr)
	if err != nil {
		u.Send(w, "Invalid reminder ID", nil, http.StatusBadRequest)
		return
	}

	listIDStr := r.PathValue("list_id")
	listID, err := strconv.Atoi(listIDStr)
	if err != nil {
		u.Send(w, "Invalid reminder ID", nil, http.StatusBadRequest)
		return
	}

	var data models.OmeletteList
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, "Invalid request body", nil, http.StatusBadRequest)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	list, err := h.Omelette.UpdateList(userID, tableID, listID, &data)
	if err != nil {
		u.Send(w, "Could not update list: "+err.Error(), nil, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", list, http.StatusOK)
}

func (h *OmeletteHandler) DeleteList(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	tableIDStr := r.PathValue("id")
	listIDStr := r.PathValue("list_id")

	tableID, err := strconv.Atoi(tableIDStr)
	if err != nil {
		u.Send(w, "Invalid event ID", nil, http.StatusBadRequest)
		return
	}

	listID, err := strconv.Atoi(listIDStr)
	if err != nil {
		u.Send(w, "Invalid reminder ID", nil, http.StatusBadRequest)
		return
	}

	err = h.Omelette.DeleteList(userID, tableID, listID)
	if err != nil {
		u.Send(w, "Could not delete list: "+err.Error(), nil, http.StatusInternalServerError)
		return
	}

	u.Send(w, "DELETED", nil, http.StatusOK)
}

func (h *OmeletteHandler) CreateCard(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	user_id, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	listIDStr := r.PathValue("list_id")
	list_id, err := strconv.Atoi(listIDStr)
	if err != nil {
		u.Send(w, "Error reading the ID requested", nil, http.StatusBadRequest)
		return
	}

	var data models.OmeletteCard
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, err.Error(), nil, http.StatusConflict)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	list, err := h.Omelette.CreateCard(user_id, list_id, &data)
	if err != nil {
		u.Send(w, "Could not create card: "+err.Error(), list, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", list, http.StatusCreated)
}

func (h *OmeletteHandler) GetCardByID(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	cardIDStr := r.PathValue("card_id")
	cardID, err := strconv.Atoi(cardIDStr)
	if err != nil {
		u.Send(w, "Invalid reminder ID", nil, http.StatusBadRequest)
		return
	}

	card, err := h.Omelette.GetCardByID(userID, cardID)
	if err != nil {
		u.Send(w, "Could not retrieve card: "+err.Error(), nil, http.StatusNotFound)
		return
	}

	u.Send(w, "", card, http.StatusOK)
}

func (h *OmeletteHandler) GetAllCards(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	cards, err := h.Omelette.GetAllCards(userID)
	if err != nil {
		u.Send(w, "Could not retrieve cards: "+err.Error(), nil, http.StatusNotFound)
		return
	}

	u.Send(w, "", cards, http.StatusOK)
}

func (h *OmeletteHandler) UpdateCard(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	cardIDStr := r.PathValue("card_id")
	cardID, err := strconv.Atoi(cardIDStr)
	if err != nil {
		u.Send(w, "Invalid card ID", nil, http.StatusBadRequest)
		return
	}

	var data models.OmeletteCard
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		u.Send(w, "Invalid request body", nil, http.StatusBadRequest)
		return
	}

	msg, err := models.ValidateModel(data)
	if err != nil {
		u.Send(w, "Invalid request", msg, http.StatusBadRequest)
		return
	}

	card, err := h.Omelette.UpdateCard(userID, cardID, &data)
	if err != nil {
		u.Send(w, "Could not update card: "+err.Error(), nil, http.StatusInternalServerError)
		return
	}

	u.Send(w, "", card, http.StatusOK)
}

func (h *OmeletteHandler) DeleteCard(w http.ResponseWriter, r *http.Request) {
	userClaims := u.GetUserFromContext(r.Context())
	if userClaims == nil {
		u.Send(w, "Error: user context is empty", nil, http.StatusInternalServerError)
		return
	}

	userID, err := u.ParseUUID(w, userClaims.ID)
	if err != nil {
		return
	}

	cardIDStr := r.PathValue("card_id")
	cardID, err := strconv.Atoi(cardIDStr)
	if err != nil {
		u.Send(w, "Invalid card ID", nil, http.StatusBadRequest)
		return
	}

	err = h.Omelette.DeleteCard(userID, cardID)
	if err != nil {
		u.Send(w, "Could not delete card: "+err.Error(), nil, http.StatusInternalServerError)
		return
	}

	u.Send(w, "DELETED", nil, http.StatusOK)
}
