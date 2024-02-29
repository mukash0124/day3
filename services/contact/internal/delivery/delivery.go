package delivery

import (
	"architecture_go/services/contact/internal/domain"
	"architecture_go/services/contact/internal/repository"
	"encoding/json"
	"net/http"
	"strconv"
)

type ContactDelivery interface {
	CreateContactHandler(w http.ResponseWriter, r *http.Request)
	ReadContactHandler(w http.ResponseWriter, r *http.Request)
	UpdateContactHandler(w http.ResponseWriter, r *http.Request)
	DeleteContactHandler(w http.ResponseWriter, r *http.Request)

	CreateGroupHandler(w http.ResponseWriter, r *http.Request)
	ReadGroupHandler(w http.ResponseWriter, r *http.Request)

	AddContactToGroupHandler(w http.ResponseWriter, r *http.Request)
}

type ContactHandler struct {
	Repo repository.ContactRepository
}

func (h *ContactHandler) ReadContactHandler(w http.ResponseWriter, r *http.Request) {
	contactID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	contact, err := h.Repo.ReadContact(contactID)
	if err != nil {
		http.Error(w, "Contact not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(contact)
}

func (h *ContactHandler) UpdateContactHandler(w http.ResponseWriter, r *http.Request) {
	contactID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	var newContact domain.Contact
	if err := json.NewDecoder(r.Body).Decode(&newContact); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Repo.UpdateContact(contactID, newContact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ContactHandler) DeleteContactHandler(w http.ResponseWriter, r *http.Request) {
	contactID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	err = h.Repo.DeleteContact(contactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ContactHandler) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	var group domain.Group
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.Repo.CreateGroup(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ContactHandler) ReadGroupHandler(w http.ResponseWriter, r *http.Request) {
	groupID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	group, err := h.Repo.ReadGroup(groupID)
	if err != nil {
		http.Error(w, "Group not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(group)
}

func (h *ContactHandler) AddContactToGroupHandler(w http.ResponseWriter, r *http.Request) {
	contactID, err := strconv.Atoi(r.URL.Query().Get("contact_id"))
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	groupID, err := strconv.Atoi(r.URL.Query().Get("group_id"))
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	err = h.Repo.AddContactToGroup(contactID, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
